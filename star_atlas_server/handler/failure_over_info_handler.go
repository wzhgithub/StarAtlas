package handler

import (
	"errors"
	"fmt"
	"os/exec"
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
func GetFailureOverInfoList(c *gin.Context) {
	fr_list, err := GetFailureOverRequestList(true)
	if err != nil {
		glog.Errorf("failed read failure over request list from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ret_fr_list := []model.FailureOverRequest{}
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

type DoFailureOverRequest struct {
	TransType  uint8  `json:"transType"`  // 整机 0 分区 1 必填
	FromVmcId  uint8  `json:"fromVmcId"`  // vmcid 必填
	IsFault    uint8  `json:"isFault"`    // [0, 1] false true
	DeviceId   uint8  `json:"deviceId"`   // 分区 必填
	TaskName   string `json:"taskName"`   // 0
	TaskType   uint8  `json:"taskType"`   // 0
	AppId      uint8  `json:"appId"`      // 必填
	ToDeviceId uint8  `json:"toDeviceId"` // 必填 转移到的设备绑卡
}

func (dfr *DoFailureOverRequest) GenPack() (string, string, error) {
	if dfr == nil {
		return "", "", errors.New("DoFailureOverRequest nil objects")
	}
	if dfr.TransType == 0 {
		return "0", fmt.Sprintf("%d|%d", dfr.FromVmcId, dfr.IsFault), nil
	}
	if dfr.TransType == 1 {
		return "1", fmt.Sprintf("%d|%d|%s|%d|%d|%d", dfr.FromVmcId, dfr.DeviceId, dfr.TaskName, dfr.TaskType, dfr.AppId, dfr.IsFault), nil
	}
	return "", "", fmt.Errorf("DoFailureOverRequest unknown req:%v", dfr)
}

func (dfr *DoFailureOverRequest) GenUniqueKey() string {
	if dfr == nil {
		return ""
	}

	return fmt.Sprintf("%d_%d_%d_%d", dfr.TransType, dfr.FromVmcId, dfr.DeviceId, dfr.ToDeviceId)
}

func (dfr *DoFailureOverRequest) ToFailureOverEntity() *model.FailureOverRequest {
	foe := &model.FailureOverRequest{
		From:        model.FailureOverInfo{},
		To:          model.FailureOverInfo{},
		TransStatus: 200,
		UniqueKey:   "",
	}

	if dfr != nil {
		foe.UniqueKey = dfr.GenUniqueKey()
		foe.From.AppID = fmt.Sprintf("%d", dfr.AppId)
		foe.From.DeviceId = fmt.Sprintf("%d", dfr.DeviceId)
		foe.From.VMCID = fmt.Sprintf("%d", dfr.FromVmcId)
		foe.To.AppID = fmt.Sprintf("%d", dfr.AppId)
		foe.To.DeviceId = fmt.Sprintf("%d", dfr.ToDeviceId)
		foe.To.VMCID = fmt.Sprintf("%d", dfr.FromVmcId)
	}

	return foe
}

func GetFailureOverInfo(ctx *gin.Context) {
	req := &DoFailureOverRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ff := bson.M{"unique_key": req.GenUniqueKey()}
	ts := bson.M{"trans_status": 200}
	filter := bson.M{"$and": []bson.M{ff, ts}}
	rsp := &model.FailureOverRequest{}
	err = mgm.CollectionByName(cFailureOverTable).First(filter, rsp)
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ctx.JSON(200, model.NewCommonResponseSucc(rsp))
}

// do failure over request
func DoFailureOver(ctx *gin.Context) {
	req := &DoFailureOverRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	a1, a2, err := req.GenPack()
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	cmd := exec.Command("./trans", a1, a2)
	out, err := cmd.Output()
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	// save from entity state
	f := req.ToFailureOverEntity()
	if err = mgm.CollectionByName(cFailureOverTable).Create(f); err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ctx.JSON(200, model.NewCommonResponseSucc(out))
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

func GetFailureOverRequestList(unfinish_req bool) ([]model.FailureOverRequest, error) {
	fr_list := []model.FailureOverRequest{}

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
