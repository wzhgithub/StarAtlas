package handler

import (
	"fmt"
	"os/exec"
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
)

type VMCDataRspJson struct {
	Success bool              `json:"success"`
	Data    model.VMCDataJson `json:"data"`
	Code    uint8             `json:"code"`
	Msg     string            `json:"msg"`
}

const cFailureOverTable = "failure_over_log"

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
		glog.Errorf("failed read vmcdata from db, error: %s\n", err.Error())
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	// glog.Infof("vmcdata_read: %+v\n", vmcdata_read)
	// vmcdata_rsp := vmcdata_read.TransferVMCDataToJson()
	// if vmcdata_rsp == nil {
	// 	glog.Errorf("failed to transfer vmcdata into Json")
	// 	c.JSON(500, model.NewCommonResponseFail(err))
	// 	return
	// }

	// construct topo
	topo := &model.TopoTable{}
	err = topo.CollectOp()
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		glog.Errorf("Failed to collect topo from db, error: %s\n", err.Error())
		return
	}
	status, err := topo.GetVmcStatus(vmc_id)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		glog.Errorf("Failed to GetVmcStatus, when vmc_id = %d, error: %s\n", vmc_id, err.Error())
		return
	}
	var s *model.CommonResponse
	if status == "RUN" {
		s = model.NewCommonResponseSucc(*vmcdata_read)
		c.JSON(200, s)
		return
	}
	s = model.NewCommonResponseSucc(status)
	c.JSON(200, s)
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
			c.JSON(500, &VMCSequenceRspJson{Success: false, Msg: "Parse vmcid error"})
			return
		}
		vmcs, err := vmcdata_read.GetVMCList(int32(vmcid))

		if err != nil && len(vmcs) == 0 {
			glog.Errorf("get vmcid: %s faild", vmcid)
			c.JSON(500, &VMCSequenceRspJson{Success: false, Msg: "Get vmcid error"})
			return
		}

		rsp := &VMCSequenceRspJson{}
		rsp.Success = true
		for _, v := range vmcs {
			var cpuComputingPower uint16
			for _, e := range v.CPUSet {
				cpuComputingPower += e.FloatComputingPower
			}
			var gpuComputingPower uint16
			for _, e := range v.GPUSet {
				gpuComputingPower += e.FloatComputingPower
			}
			var dspComputingPower uint16
			for _, e := range v.DSPSet {
				dspComputingPower += e.IntComputingPower
			}
			vmc_status := &model.VMCStatus{
				UpdatedAt:            v.UpdatedAt,
				CPUComputingPower:    cpuComputingPower,
				GPUComputingPower:    gpuComputingPower,
				DSPIntComputingPower: dspComputingPower,
				MomoryUsage:          v.MemoryUsage,
				DiskUsage:            v.TotalDiskUsage,
				TotalUsage:           v.TotalCPUUsage,
				GpuUsage:             v.TotalGPUUsage,
				DspUsage:             v.TotalDSPUsage,
			}
			rsp.Data = append(rsp.Data, *vmc_status)
		}
		rsp.Code = 0
		rsp.Msg = "ok"

		c.JSON(200, rsp)

	} else {
		c.JSON(500, &VMCSequenceRspJson{Success: false, Msg: "request without vmcid"})
		return
	}
}

func FailureOver(c *gin.Context) {
	req := &model.FailureOverRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	t := &model.TopoTable{}
	if err := t.CollectOp(); err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	vid, err := strconv.ParseInt(req.From.VMCID, 10, 64)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	bids, err := t.GetBackupId(vid)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	req.To.VMCID = fmt.Sprintf("%d", bids[0])
	glog.Infof("failure request %+v\n", req)
	mockBin := "./trans"
	arg1 := req.From.VMCID
	arg2 := req.To.VMCID
	cmd := exec.Command(mockBin, arg1, arg2)
	stdout, err := cmd.Output()
	req.TransStatus = 500 // unfinished
	req.UniqueKey = fmt.Sprintf("%s_%s_%s_%s_%s_%s",
		req.From.VMCID, req.From.AppID, req.From.TaskID, req.To.VMCID, req.To.AppID, req.To.TaskID)
	if err != nil {
		glog.Errorf("run command:%+v failed err:%s\n", cmd, err.Error())
	}
	glog.Infof("cmd output %s\n", stdout)
	if err = mgm.CollectionByName(cFailureOverTable).Create(req); err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}

	c.JSON(200, model.NewCommonResponseSucc(""))
}
