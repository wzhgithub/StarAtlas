package handler

import (
	"fmt"
	"handler"
	"star_atlas_server/model"

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
func GetFailureOverInfo(c *gin.Context) {
	fr_list, err := GetFailureOverRequestList()
	if err != nil {
		glog.Error("failed read failure over request list from db, error: %s\n", err.Error())
		c.JSON(400, model.NewCommonResponseFail(err))
		return
	}

	c.JSON(200, model.NewCommonResponseSucc(fr_list[0]))

}

func GetFailureOverVMCData(c *gin.Context) {
	fr_list, err := GetFailureOverRequestList()
	if err != nil {
		glog.Error("failed read failure over request list from db, error: %s\n", err.Error())
		c.JSON(400, model.NewCommonResponseFail(err))
		return
	}
	fr := fr_list[0]
	from_vmc_id, to_vmc_id := fr.From.VMCID, fr.To.VMCID
	from_vmcdata, to_vmcdata := &model.VMCData{}, &model.VMCData{}
	err = from_vmcdata.CollectVMCData(int32(from_vmc_id))
	if err != nil {
		glog.Error("failed read vmcdata from db, error: %s\n", err.Error())
		c.JSON(400, model.NewCommonResponseFail(err))
		return
	}

	err = to_vmcdata.CollectVMCData(int32(to_vmc_id))
	if err == nil {
		// TODO:change status
	}

	failure_over_vmc_data := FailureOverVMCData{}
	// TODO: access topo interface
	// check from_vmc_id valid
	// check to_vmc_id valid
	c.JSON(200, model.NewCommonResponseSucc(failure_over_vmc_data))

}

type FailureOverVMCData struct {
	From model.VMCData `json:"from"`
	To   model.VMCData `json:"to"`
}

func GetFailureOverRequestList() ([]handler.FailureOverRequest, error) {
	fr_list := []handler.FailureOverRequest{}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	ret := mgm.CollectionByName("failure_over_log").SimpleFind(&fr_list, bson.M{"type": "vmc"}, findOptions)
	if len(fr_list) < 1 {
		return fr_list, fmt.Errorf("fr_list is empty")
	}
	return fr_list, ret
}
