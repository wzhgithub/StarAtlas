package handler

import (
	"net/http"
	"star_atlas_server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// show
func TopoTable(c *gin.Context) {
	topo := &model.TopoTable{}
	err := topo.CollectOp()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Failed to collect topo from db",
		})
		glog.Error("Failed to collect topo from db, error: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": topo,
	})
}

// insert
func Insert(c *gin.Context) {
	topo := &model.TopoTable{}
	name := c.Request.FormValue("name")
	deviceType := c.Request.FormValue("deviceType")
	deviceStatus := c.Request.FormValue("deviceStatus")
	parentId, err := strconv.ParseUint(c.Request.FormValue("parentId"), 10, 16)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Parse parentId failed",
		})
		glog.Error("Parse parentId failed")
		return
	}
	upstreamId, err := strconv.ParseUint(c.Request.FormValue("upstreamId"), 10, 16)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Parse upstreamId failed",
		})
		glog.Error("Parse upstreamId failed")
		return
	}
	node := &model.Nodes{
		Id:           0,
		Name:         name,
		DeviceType:   deviceType,
		ParentId:     uint16(parentId),
		UpstreamId:   uint16(upstreamId),
		DeviceStatus: deviceStatus,
		OtherInfo:    nil,
	}
	err = topo.InsertOp(node)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "InsertOp failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Insert success",
		"data": node,
	})
}

// modify
func Update(c *gin.Context) {

}

// delete
func Delete(c *gin.Context) {
	topo := &model.TopoTable{}
	deviceId, err := strconv.ParseUint(c.Request.FormValue("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invaild id",
		})
		glog.Error("Invaild id")
		return
	}
	err = topo.DeleteOp(uint16(deviceId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "DeleteOp failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Delete success",
		"data": deviceId,
	})
}
