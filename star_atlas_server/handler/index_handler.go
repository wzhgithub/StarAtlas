package handler

import (
	"fmt"
	"net/http"
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
func Index(c *gin.Context) {
	c.JSON(200, "ok")
}

func VMCStatusTest(c *gin.Context) {
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

	for _, node := range topo.Node {
		if node.DeviceType == "vmc" {
			status, err := topo.GetVmcStatus(node.Id)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  fmt.Sprintf("Failed to GetVmcStatus, when vmc_id = %d", node.Id),
				})
				glog.Errorf("Failed to GetVmcStatus, when vmc_id = %d, error: %s\n", node.Id, err.Error())
			}
			glog.Infof("vmc_id: %d, status: %s", node.Id, status)

			backupId, err1 := topo.GetBackupId(node.Id)
			if err1 != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  fmt.Sprintf("Failed to GetBackupId, when vmc_id = %d", node.Id),
				})
				glog.Errorf("Failed to GetBackupId, when vmc_id = %d, error: %s\n", node.Id, err.Error())
			}
			glog.Infof("vmc_id: %d, backupId: %+v", node.Id, backupId)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
