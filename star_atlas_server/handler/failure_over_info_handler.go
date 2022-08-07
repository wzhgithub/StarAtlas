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

// reture all of unfinished failure over requests
func GetFailureOverInfo(c *gin.Context) {
	fr_list, err := GetFailureOverRequestList(true)
	if err != nil {
		glog.Errorf("failed read failure over request list from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ret_fr_list := []FailureOverRequest{}
	filter_map := make(map[string]bool)
	for _, fr := range fr_list {
		_, ok := filter_map[fr.UniqueKey]
		if ok {
			continue
		}
		filter_map[fr.UniqueKey] = true
		ret_fr_list = append(ret_fr_list, fr)
	}
	c.JSON(200, model.NewCommonResponseSucc(ret_fr_list))
}

// reture the most recently failure over vmcdata
func GetFailureOverVMCData(c *gin.Context) {
	fr_list, err := GetFailureOverRequestList(false)
	if err != nil {
		glog.Errorf("failed read failure over request list from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	fr := fr_list[0] // most recent failure over request
	from_vmc_id, _ := strconv.ParseInt(fr.From.VMCID, 10, 32)
	to_vmc_id, _ := strconv.ParseInt(fr.To.VMCID, 10, 32)

	from_vmcdata, to_vmcdata := &model.VMCData{}, &model.VMCData{}
	err = from_vmcdata.CollectVMCData(int32(from_vmc_id))
	if err != nil {
		glog.Errorf("failed read vmcdata from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	err = to_vmcdata.CollectVMCData(int32(to_vmc_id))
	if err != nil {
		glog.Errorf("failed read vmcdata from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	// if to_vmcdata tasks are not empty
	if len(to_vmcdata.APPInfo) > 0 && len(to_vmcdata.APPInfo[0].TaskSet) > 0 {
		glog.Info("failure over request from %d to %d finished!\n", from_vmc_id, to_vmc_id)
		// see https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-update-documents
		// see https://www.cnblogs.com/williamjie/p/9598748.html
		ff := bson.M{"from": bson.M{"vmc_id": fmt.Sprintf("%d", from_vmc_id)}}
		tt := bson.M{"to": bson.M{"vmc_id": fmt.Sprintf("%d", to_vmc_id)}}
		ts := bson.M{"trans_status": 500}
		filter := bson.M{"$and": []bson.M{ff, tt, ts}}
		updateM := bson.M{"$set": bson.M{"trans_status": 200}}
		res, err := mgm.CollectionByName(cFailureOverTable).UpdateOne(mgm.Ctx(), filter, updateM)
		if err != nil {
			glog.Errorf("update failed filter:%+v update:%+v err: %v\n", filter, updateM, err)
			c.JSON(500, model.NewCommonResponseFail(err))
			return
		}
		glog.Infof("update succeeded result: %+v\n", res)
	}
	failure_over_vmc_data := FailureOverVMCData{}
	failure_over_vmc_data.From = *from_vmcdata
	failure_over_vmc_data.To = *to_vmcdata
	c.JSON(200, model.NewCommonResponseSucc(failure_over_vmc_data))

}

type FailureOverVMCData struct {
	From model.VMCData `json:"from"`
	To   model.VMCData `json:"to"`
}

func GetFailureOverRequestList(unfinish_req bool) ([]FailureOverRequest, error) {
	fr_list := []FailureOverRequest{}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	filter := bson.M{}
	if unfinish_req {
		filter = bson.M{"trans_status": 500}
	}
	ret := mgm.CollectionByName(cFailureOverTable).SimpleFind(&fr_list, filter, findOptions)
	if len(fr_list) < 1 {
		return fr_list, fmt.Errorf("fr_list is empty")
	}
	return fr_list, ret
}
