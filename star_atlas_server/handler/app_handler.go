package handler

import (
	"net/http"
	"star_atlas_server/config"
	"star_atlas_server/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	coll.SimpleFind(&vmcs, bson.M{"vmc_id": v.VmcId}, findOptions)
	if len(vmcs) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "vmc_id is not exist in db",
		})
		glog.Errorf("vmc_id is not exist in db")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "get app success",
		"data": vmcs[0].APPInfo,
	})
}
