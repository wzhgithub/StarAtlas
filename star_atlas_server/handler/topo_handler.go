package handler

import (
	"net/http"
	"star_atlas_server/model"

	"github.com/gin-gonic/gin"
)

func ShowTopo(c *gin.Context) {
	var topo model.TopoTable
	topo.CollectOp()
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
	})
}
