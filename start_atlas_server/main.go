package main

import (
	"flag"
	"fmt"
	"log"
	"start_atlas_server/config"
	"start_atlas_server/db"
	"start_atlas_server/handler"
	"time"

	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "path", "./config.yaml", "yaml文件加载路径")
}

func main() {
	flag.Parse()
	err := config.Init(configPath)
	log.Println("c", config.CommonConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	go handler.UdpDataRev(config.CommonConfig.UDPPort)
	go handler.ParseData()
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}), gin.Recovery())

	router.GET("/test", handler.Index)
	router.Run(":9999")

}
