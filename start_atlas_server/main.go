package main

import (
	"flag"
	"fmt"
	_ "start_atlas_server/db"
	"start_atlas_server/handler"
	"time"

	"github.com/gin-gonic/gin"
)

var udpPort int

func init() {
	flag.IntVar(&udpPort, "port", 8080, "Port to listen on")
}

func main() {
	flag.Parse()
	go handler.UdpDataRev(udpPort)
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
