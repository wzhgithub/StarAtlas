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

type details struct {
	dbase  int64
	dtypes int
	dcores int
}

var deviceDetails = map[string]details{
	"CPU":  {model.C_CPU_BASE, 5, 256},
	"GPU":  {model.C_GPU_BASE, 1, 256},
	"DSP":  {model.C_DSP_BASE, 3, 256},
	"FPGA": {model.C_FPGA_BASE, 1, 1},
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
	if _, ok := deviceDetails[dType]; !ok {
		return fmt.Errorf("cannot find device type: %s", dType)
	}
	n := &model.Nodes{
		Id:           vmcId*model.CVMCBase + deviceDetails[dType].dbase,
		Name:         dType + "_all",
		DeviceType:   strings.ToLower(dType),
		ParentId:     uint16(vmcId),
		UpstreamId:   0,
		DeviceStatus: "RUN",
		DeviceNum:    int32(num),
		OtherInfo:    make([]*model.OtherInfos, 0),
	}
	device_ids := make([]string, 0)
	device_names := make([]string, 0)
	device_types := make([]string, 0)
	device_cores := make([]string, 0)
	for i := 0; i < num; i++ {
		cur_id := vmcId*model.CVMCBase + int64(i)
		device_ids = append(device_ids, fmt.Sprintf("%d", cur_id))
		device_names = append(device_names, dType+"_"+strconv.Itoa(i))
		device_types = append(device_types, fmt.Sprintf("%d", rand.Intn(deviceDetails[dType].dtypes)))
		device_cores = append(device_cores, fmt.Sprintf("%d", rand.Intn(deviceDetails[dType].dcores)))
	}
	n.OtherInfo = append(n.OtherInfo, model.NewOtherInfos("cpu_ids", device_ids))
	n.OtherInfo = append(n.OtherInfo, model.NewOtherInfos("cpu_names", device_names))
	n.OtherInfo = append(n.OtherInfo, model.NewOtherInfos("cpu_types", device_types))
	n.OtherInfo = append(n.OtherInfo, model.NewOtherInfos("cpu_cores", device_cores))

	if err := topo.InsertOp(n); err != nil {
		return fmt.Errorf("insert %s device failed", dType)
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
