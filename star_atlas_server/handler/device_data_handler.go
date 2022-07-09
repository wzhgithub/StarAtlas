package handler

import (
	"star_atlas_server/model"
	"strconv"

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
func GetDeviceData(c *gin.Context) {
	vmc_id_64, _ := strconv.ParseInt(c.Query("vmc_id"), 10, 64)
	vmc_id := int32(vmc_id_64)
	device_type := c.Query("device_type")

	device_data_list, err := model.CollectDeviceData(vmc_id, device_type)
	if err != nil {
		glog.Error("failed read device data from db, error: %s\n", err.Error())
		c.JSON(400, &DeviceDataRspJson{Success: false, Msg: "fail"})
		return
	}

	rsp := &DeviceDataRspJson{}
	rsp.Success = true
	for _, v := range device_data_list {
		rsp.Data = append(rsp.Data, *v)
	}
	rsp.Code = 0
	rsp.Msg = "ok"

	c.JSON(200, rsp)

}

type DeviceDataRspJson struct {
	Success bool               `json:"success"`
	Data    []model.DeviceData `json:"data"`
	Code    uint8              `json:"code"`
	Msg     string             `json:"msg"`
}
