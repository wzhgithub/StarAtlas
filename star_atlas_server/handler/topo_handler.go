package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"star_atlas_server/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var deviceType = map[string]bool{
	"CPU":  true,
	"GPU":  true,
	"DSP":  true,
	"FPGA": true,
}

const cRange = 10

// show
func TopoShow(c *gin.Context) {
	topo := &model.TopoTable{}
	err := topo.CollectOp()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to collect topo from db",
		})
		glog.Errorf("Failed to collect topo from db, error: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": topo,
	})
}

// insert
func TopoInsert(c *gin.Context) {
	topo := &model.TopoTable{}
	maxId, err := topo.GetMaxVmcId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		glog.Errorf(err.Error())
		return
	}

	node := &model.Nodes{}
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "parse json failed",
		})
		return
	}
	node.Id = maxId + 1
	node.DeviceStatus = "RUN"
	node.DeviceType = "vmc"
	if err := topo.InsertOp(node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "InsertOp failed",
		})
		return
	}

	cpuNum := rand.Intn(cRange)
	gpuNum := rand.Intn(cRange)
	fpgaNum := rand.Intn(cRange)
	dspNum := rand.Intn(cRange)
	if err := insertDevice(topo, cpuNum, "CPU", node.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Insert cpu device failed",
		})
		glog.Errorf(err.Error())
		return
	}
	if err := insertDevice(topo, gpuNum, "GPU", node.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Insert gpu device failed",
		})
		glog.Errorf(err.Error())
		return
	}
	if err := insertDevice(topo, fpgaNum, "FPGA", node.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Insert fpga device failed",
		})
		glog.Errorf(err.Error())
		return
	}
	if err := insertDevice(topo, dspNum, "DSP", node.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Insert dsp device failed",
		})
		glog.Errorf(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Insert success",
		"data": node,
	})
}

func insertDevice(topo *model.TopoTable, num int, dType string, vmcId int64) error {
	dType = strings.ToUpper(dType)
	if !deviceType[dType] {
		return fmt.Errorf("cannot find device type: %s", dType)
	}
	for i := 0; i < num; i++ {
		n := &model.Nodes{
			Id:           vmcId*model.CVMCBase + int64(i),
			Name:         dType + "_" + strconv.Itoa(i),
			DeviceType:   strings.ToLower(dType),
			ParentId:     uint16(vmcId),
			UpstreamId:   0,
			DeviceStatus: "RUN",
			OtherInfo:    nil,
		}
		if err := topo.InsertOp(n); err != nil {
			return fmt.Errorf("insert %s device failed", dType)
		}
	}
	return nil
}

// delete
func TopoDelete(c *gin.Context) {
	topo := &model.TopoTable{}
	if err := topo.CollectOp(); err != nil {
		glog.Errorln("get topo table failed")
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}

	node := &model.Nodes{}
	if err := c.ShouldBindJSON(&node); err != nil {
		glog.Errorln("Invaild id")
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}

	delNodes := make([]int64, 0)
	for _, tNode := range topo.Node {
		if tNode.ParentId == uint16(node.Id) || (tNode.Id == node.Id && tNode.DeviceType == "vmc") {
			glog.Infof("Delete vmc node %+v\n", tNode)
			delNodes = append(delNodes, tNode.Id)
		}
	}
	if err := topo.DeleteOp(delNodes); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}
	c.JSON(http.StatusOK, model.NewCommonResponseSucc(node.Id))
}
