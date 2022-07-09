package handler

import (
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
func GetVMCData(c *gin.Context) {
	vmc_id_64, _ := strconv.ParseInt(c.Query("vmc_id"), 10, 64)
	vmc_id := int32(vmc_id_64)
	vmcdata_read := &model.VMCData{}
	err := vmcdata_read.CollectVMCData(vmc_id)
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

type VMCDataRspJson struct {
	Success bool              `json:"success"`
	Data    model.VMCDataJson `json:"data"`
	Code    uint8             `json:"code"`
	Msg     string            `json:"msg"`
}

func GetVMCSequence(c *gin.Context) {
	paramPairs := c.Request.URL.Query()
	vid, ok := paramPairs["vmcid"]
	vmcdata_read := &model.VMCData{}
	if ok && len(vid) > 0 {
		vmcdata_read.GetVMCList(vid[0])
	} else {
		c.JSON(400, &RspJson{Success: false, Msg: "request without vmcid"})
		return
	}
}
