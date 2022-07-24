package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"star_atlas_server/model"
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
	if err := topo.InsertOp(node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "InsertOp failed",
		})
		return
	}

	cpuNum := rand.Intn(256)
	gpuNum := rand.Intn(256)
	fpgaNum := rand.Intn(256)
	dspNum := rand.Intn(256)
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
			Name:         dType + "_" + string(rune(i)),
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
	node := &model.Nodes{}
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invaild id",
		})
		glog.Errorf("Invaild id")
		return
	}
	if node.DeviceType != "vmc" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "This type of device is not vmc",
		})
		glog.Errorf("This type of device is not vmc")
		return
	}
	for _, tNode := range topo.Node {
		if tNode.ParentId == uint16(node.Id) || tNode.Id == node.Id {
			if err := topo.DeleteOp(node.Id); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "DeleteOp failed",
				})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Delete success",
		"data": node.Id,
	})
}
