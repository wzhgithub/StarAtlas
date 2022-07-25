package handler

import (
	"fmt"
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func GetAppInfo(c *gin.Context) {
	vmc_id_64, _ := strconv.ParseInt(c.Query("vmc_id"), 10, 64)
	vmc_id := int32(vmc_id_64)

	// collect app_info_list
	app_info_list, err := model.CollectAppInfo(vmc_id)
	if err != nil {
		glog.Errorf("failed read app info from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	// construct topo
	topo := &model.TopoTable{}
	if err := topo.CollectOp(); err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		glog.Errorf("Failed to collect topo from db, error: %s\n", err.Error())
		return
	}
	status, err := topo.GetVmcStatus(vmc_id_64)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		glog.Errorf("Failed to GetVmcStatus, when vmc_id = %d, error: %s\n", vmc_id, err.Error())
		return
	}
	if status != "RUN" {
		c.JSON(500, model.NewCommonResponseSucc(status))
		glog.Errorf("vmc_id = %d, status is: %s\n", vmc_id, status)
		return

	}

	// check failure over database
	fr_list := []FailureOverRequest{}
	apps := []model.App{}
	app_info_rsp := AppInfoRsp{IsTransfer: false, Apps: apps}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	ff := bson.M{"from": bson.M{"vmc_id": fmt.Sprintf("%d", vmc_id_64)}}
	ts := bson.M{"trans_status": 500}
	filter := bson.M{"$and": []bson.M{ff, ts}}

	err = mgm.CollectionByName(cFailureOverTable).SimpleFind(&fr_list, filter, findOptions)
	if err != nil {
		glog.Errorf("failed read trans record from failure over request list db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	if len(fr_list) > 0 {
		app_info_rsp.IsTransfer = true
		fr := fr_list[0]
		for _, app := range app_info_list {
			from_app_id_uint64, _ := strconv.ParseInt(fr.From.AppID, 10, 8)
			from_app_id := uint8(from_app_id_uint64)
			if from_app_id == app.ID {
				app.IsTransfer = true
			}
		}
	}

	for _, app_info := range app_info_list {
		app_info_rsp.Apps = append(app_info_rsp.Apps, *app_info)
	}

	c.JSON(200, model.NewCommonResponseSucc(app_info_rsp))

}

type AppInfoRsp struct {
	IsTransfer bool        `json:"is_transfer"`
	Apps       []model.App `json:"apps"`
}
