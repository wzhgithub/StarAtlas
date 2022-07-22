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
func GetAppInfo(c *gin.Context) {
	vmc_id_64, _ := strconv.ParseInt(c.Query("vmc_id"), 10, 64)
	vmc_id := int32(vmc_id_64)

	app_info_list, err := model.CollectAppInfo(vmc_id)
	glog.Error("app info list: %+v\n", app_info_list)
	if err != nil {
		glog.Error("failed read app info from db, error: %s\n", err.Error())
		c.JSON(400, model.NewCommonResponseFail(err))
		return
	}

	var apps []model.App
	for _, v := range app_info_list {
		apps = append(apps, *v)
	}
	c.JSON(200, model.NewCommonResponseSucc(apps))

}
