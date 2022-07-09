package handler

import (
	"net/http"
	"star_atlas_server/config"
	"star_atlas_server/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type vmc struct {
	VmcId int64 `json:"vmc_id" bson:"vmc_id"`
}

func AppShow(c *gin.Context) {
	vmcs := []model.VMCData{}
	v := &vmc{}
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Invaild id",
		})
		glog.Errorf("Invaild id")
		return
	}

	coll := mgm.CollectionByName(config.CommonConfig.DBVMCDataTableName)
	coll.SimpleFind(&vmcs, bson.M{"vmc_id": v.VmcId})
	if len(vmcs) != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Id is repeated in the db",
		})
		glog.Errorf("Id is repeated in the db")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "get app success",
		"data": vmcs[0].APPInfo,
	})
}
