package handler

import (
	"star_atlas_server/model"

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
	vmcdata_read := &model.VMCData{}
	err := vmcdata_read.CollectVMCData()
	if err != nil {
		glog.Error("failed read vmcdata from db, error: %s\n", err.Error())
	}
	glog.Infof("vmcdata_read: %+v\n", vmcdata_read)

	vmcdata_rsp := &VMCDataJson{}
	TransferVMCDataToJson(vmcdata_rsp, vmcdata_read)

	rsp := &RspJson{}
	rsp.Success = true
	rsp.Data = *vmcdata_rsp
	rsp.Code = 0
	rsp.Msg = "ok"

	c.JSON(200, rsp)

}

type RspJson struct {
	Success bool        `json:"success" bson:"success"`
	Data    VMCDataJson `json:"data"`
	Code    uint8       `json:"code"`
	Msg     string      `json:"msg"`
}

type VMCDataJson struct {
	VMCName        string `json:"vmc_name" bson:"vmc_name"` // 10bytes
	VMCID          uint8  `json:"vmc_id" bson:"vmc_id"`
	CPUNumber      uint8  `json:"cpu_number" bson:"cpu_number"`
	DSPNumber      uint8  `json:"dsp_number" bson:"dsp_number"`
	GPUNumber      uint8  `json:"gpu_number" bson:"gpu_number"`
	FPGANumber     uint8  `json:"fpga_number" bson:"fpga_number"`
	SwitchID       uint8  `json:"switch_id" bson:"switch_id"`
	TotalMemory    uint16 `json:"total_memory" bson:"total_memory"`
	TotalDisk      uint16 `json:"total_disk" bson:"total_disk"`
	MemoryUsage    uint8  `json:"memory_usage" bson:"memory_usage"`
	TotalCPUUsage  uint8  `json:"total_cpu_usage" bson:"total_cpu_usage"`
	TotalDSPUsage  uint8  `json:"total_dsp_usage" bson:"total_dsp_usage"`
	TotalGPUUsage  uint8  `json:"total_gpu_usage" bson:"total_gpu_usage"`
	TotalDiskUsage uint8  `json:"total_disk_usage" bson:"total_disk_usage"`
}

func TransferVMCDataToJson(dst *VMCDataJson, src *model.VMCData) {
	dst.VMCName = src.VMCName
	dst.VMCID = src.VMCID
	dst.CPUNumber = src.CPUNumber
	dst.DSPNumber = src.DSPNumber
	dst.GPUNumber = src.GPUNumber
	dst.FPGANumber = src.FPGANumber
	dst.SwitchID = src.SwitchID
	dst.TotalMemory = src.TotalMemory
	dst.TotalDisk = src.TotalDisk
	dst.MemoryUsage = src.MemoryUsage
	dst.TotalCPUUsage = src.TotalCPUUsage
	dst.TotalDSPUsage = src.TotalDSPUsage
	dst.TotalGPUUsage = src.TotalGPUUsage
	dst.TotalDiskUsage = src.TotalDiskUsage
}
