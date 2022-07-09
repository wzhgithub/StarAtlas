package handler

import (
	"net/http"
	"star_atlas_server/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// show
func TopoShow(c *gin.Context) {
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
func TopoInsert(c *gin.Context) {
	topo := &model.TopoTable{}
	node := &model.Nodes{}
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "parse json failed",
		})
		return
	}
	node.Id = time.Now().Unix()
	node.DeviceStatus = "START"
	if err := topo.InsertOp(node); err != nil {
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

// delete
func TopoDelete(c *gin.Context) {
	topo := &model.TopoTable{}
	node := &model.Nodes{}
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invaild id",
		})
		glog.Error("Invaild id")
		return
	}
	if err := topo.DeleteOp(node.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "DeleteOp failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Delete success",
		"data": node.Id,
	})
}
