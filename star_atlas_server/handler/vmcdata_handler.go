package handler

import (
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

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

type VMCSequenceRspJson struct {
	Success bool              `json:"success"`
	Data    []model.VMCStatus `json:"data"`
	Code    uint8             `json:"code"`
	Msg     string            `json:"msg"`
}

func GetVMCSequence(c *gin.Context) {
	paramPairs := c.Request.URL.Query()
	vid, ok := paramPairs["vmc_id"]
	vmcdata_read := &model.VMCData{}
	if ok && len(vid) > 0 {
		vmcid, err := strconv.ParseInt(vid[0], 10, 32)
		if err != nil {
			glog.Errorf("Parse vmcid: %s faild", vmcid)
			c.JSON(400, &VMCSequenceRspJson{Success: false, Msg: "Parse vmcid error"})
			return
		}
		vmcs, err := vmcdata_read.GetVMCList(int32(vmcid))

		if err != nil && len(vmcs) == 0 {
			glog.Errorf("get vmcid: %s faild", vmcid)
			c.JSON(400, &VMCSequenceRspJson{Success: false, Msg: "Get vmcid error"})
			return
		}

		rsp := &VMCSequenceRspJson{}
		rsp.Success = true
		for _, v := range vmcs {
			cpuComputingPower := int16(0)
			for _, e := range v.CPUSet {
				cpuComputingPower += int16(e.FloatComputingPower)
			}
			gpuComputingPower := int16(0)
			for _, e := range v.GPUSet {
				gpuComputingPower += int16(e.FloatComputingPower)
			}
			dspComputingPower := int16(0)
			for _, e := range v.DSPSet {
				dspComputingPower += int16(e.IntComputingPower)
			}
			vmc_status := &model.VMCStatus{
				UpdatedAt:            v.UpdatedAt,
				CPUComputingPower:    int16(cpuComputingPower),
				GPUComputingPower:    int16(gpuComputingPower),
				DSPIntComputingPower: int16(dspComputingPower),
				MomoryUsage:          int8(v.MemoryUsage),
				DiskUsage:            int8(v.TotalDiskUsage),
				TotalUsage:           int8(v.TotalCPUUsage),
			}
			rsp.Data = append(rsp.Data, *vmc_status)
		}
		rsp.Code = 0
		rsp.Msg = "ok"

		c.JSON(200, rsp)

	} else {
		c.JSON(400, &VMCSequenceRspJson{Success: false, Msg: "request without vmcid"})
		return
	}
}
