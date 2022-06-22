package main

import (
	"flag"
	"start_atlas_server/handler"

	"github.com/gin-gonic/gin"
)

var udpPort int

func init() {
	flag.IntVar(&udpPort, "port", 8080, "Port to listen on")
}

func main() {
	flag.Parse()
	go handler.UdpServer(udpPort)
	r := gin.Default()
	r.GET("/test", handler.Index)
	r.Run(":9999")

}
