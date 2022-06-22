package main

import (
	"github.com/gin-gonic/gin"
	"start_atlas_server/handler"
)

func main() {
	r := gin.Default()
	r.GET("/test", handler.Index)
	r.Run(":9999")

}
