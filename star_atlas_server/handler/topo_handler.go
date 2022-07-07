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
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": topo,
	})
}

// insert
func Insert(c *gin.Context) {

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
