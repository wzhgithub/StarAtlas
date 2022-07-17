package handler

import (
	"os/exec"
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
)

type FailureOverInfo struct {
	VimdID string `json:"vimd_id" bson:"vimd_id"`
	AppID  string `json:"app_id" bson:"app_id"`
	TaskID string `json:"task_id" bson:"task_id"`
}

type FailureOverRequest struct {
	mgm.DefaultModel `bson:",inline"`
	From             FailureOverInfo `json:"from" bson:"from"`
	To               FailureOverInfo `json:"to" bson:"to"`
}

type VMCDataRspJson struct {
	Success bool              `json:"success"`
	Data    model.VMCDataJson `json:"data"`
	Code    uint8             `json:"code"`
	Msg     string            `json:"msg"`
}

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
func GetVMCData(c *gin.Context) {
	vmc_id, _ := strconv.ParseInt(c.Query("vmc_id"), 10, 32)
	vmcdata_read := &model.VMCData{}
	err := vmcdata_read.CollectVMCData(int32(vmc_id))
	if err != nil {
		glog.Error("failed read vmcdata from db, error: %s\n", err.Error())
		c.JSON(400, &VMCDataRspJson{Success: false, Msg: "fail"})
		return
	}
	// glog.Infof("vmcdata_read: %+v\n", vmcdata_read)
	vmcdata_rsp := vmcdata_read.TransferVMCDataToJson()
	if vmcdata_rsp == nil {
		glog.Error("failed to transfer vmcdata into Json")
		c.JSON(400, &VMCDataRspJson{Success: false, Msg: "fail"})
		return
	}

	rsp := &VMCDataRspJson{}
	rsp.Success = true
	rsp.Data = *vmcdata_rsp
	rsp.Code = 0
	rsp.Msg = "ok"

	c.JSON(200, rsp)

}

func GetVMCSequence(c *gin.Context) {
	paramPairs := c.Request.URL.Query()
	vid, ok := paramPairs["vmc_id"]
	vmcdata_read := &model.VMCData{}
	if ok && len(vid) > 0 {
		vmcid, err := strconv.ParseInt(vid[0], 10, 32)
		if err != nil {
			glog.Errorf("Parse vmcid: %s faild", vmcid)
		}
		vmcs, err := vmcdata_read.GetVMCList(int32(vmcid))

		if err != nil && len(vmcs) == 0 {
			glog.Errorf("get vmcid: %s faild", vmcid)
		}

	} else {
		c.JSON(400, &VMCDataRspJson{Success: false, Msg: "request without vmcid"})
		return
	}
}

func FailureOver(c *gin.Context) {
	req := &FailureOverRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, model.NewCommonResponseFail(err))
		return
	}

	mockBin := ""
	arg1 := ""
	arg2 := ""
	arg3 := ""
	cmd := exec.Command(mockBin, arg1, arg2, arg3)
	stdout, err := cmd.Output()
	if err != nil {
		glog.Errorf("run command:%+v failed err:%s\n", cmd, err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	glog.Infof("cmd output %s\n", stdout)
	if err = mgm.CollectionByName("failure_over_log").Create(req); err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	c.JSON(200, model.NewCommonResponseSucc(""))
}
