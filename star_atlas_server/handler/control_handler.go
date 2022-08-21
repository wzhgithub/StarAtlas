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

const CDbSenderName = "sender_db"

type sendAddr struct {
	mgm.DefaultModel `json:",inline" bson:",inline"`
	DbName           string `json:"name" bson:"name"`
	SendIP           string `json:"send_ip" bson:"send_ip"`
	SendPort         int    `json:"send_port" bson:"send_port"`
	SendMAC          string `json:"send_mac" bson:"send_mac"`
}

func ControlSender(c *gin.Context) {
	address := &sendAddr{}
	address.DbName = CDbSenderName
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}
	if err := address.CreateOp(); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}
	c.JSON(http.StatusOK, model.NewCommonResponseSucc(address))
}

func GetControlMsg(c *gin.Context) {
	address := &sendAddr{}
	if err := address.CollectOp(); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewCommonResponseFail(err))
		return
	}
	c.JSON(http.StatusOK, model.NewCommonResponseSucc(address))
}

func (s *sendAddr) CreateOp() error {
	oldAddress := &sendAddr{}
	err := mgm.CollectionByName(config.CommonConfig.DBSenderTableName).First(bson.M{"name": CDbSenderName}, oldAddress)
	if err != nil {
		glog.Infof("[CreateOp] Cannot find, create a new sender_table")
	}
	glog.Infof("[CreateOp] sender db name = %+v", s)
	if oldAddress.DbName == "" {
		glog.Infof("[CreateOp] new sender_table: %+v", s)
		return mgm.CollectionByName(config.CommonConfig.DBSenderTableName).Create(s)
	}
	oldAddress.SendIP = s.SendIP
	oldAddress.SendPort = s.SendPort
	oldAddress.SendMAC = s.SendMAC
	return oldAddress.UpdateOp()
}

func (s *sendAddr) UpdateOp() error {
	return mgm.CollectionByName(config.CommonConfig.DBSenderTableName).Update(s)
}

func (s *sendAddr) CollectOp() error {
	return mgm.CollectionByName(config.CommonConfig.DBSenderTableName).First(bson.M{}, s)
}
