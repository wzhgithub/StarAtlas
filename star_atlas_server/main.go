package main

import (
	"flag"
	"fmt"
	"star_atlas_server/config"
	"star_atlas_server/db"
	"star_atlas_server/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nl8590687/asrt-sdk-go/sdk"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "path", "./config.yaml", "yaml文件加载路径")
}

func testWav() {
	filename := "test/record.wav"
	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		fmt.Println(err)
	}
	res, err := handler.RecogniteByType(wave.GetRawSamples(), 16000, 1, 2, "speech")
	if err != nil {
		fmt.Println(err)
	}
	glog.Infof("received speech from server %v\n", res)
	if res.StatusCode != 200 {
		fmt.Println(res.StatucMesaage)
	}
	glog.Infof("received speech word %v\n", res.Result)
}

func main() {
	flag.Parse()
	err := config.Init(configPath)
	glog.Infof("config:%+v\n", config.CommonConfig)
	if err != nil {
		glog.Fatalf(err.Error())
	}
	err = db.Init()
	if err != nil {
		glog.Fatal(err)
	}
	go handler.UdpDataRev(config.CommonConfig.UDPPort)
	go handler.ParseData()
	go handler.SatelliteTCPHandlerInit(config.CommonConfig.SatelliteTCPPort)
	testWav()
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
	router.GET("/test/vmcstatus", handler.VMCStatusTest)
	router.GET("/vmcdata", handler.GetVMCData)
	router.GET("/devicedata", handler.GetDeviceData) // usage: http://localhost:9999/devicedata?vmc_id=0&device_type=cpu
	router.GET("/appinfo", handler.GetAppInfo)       // usage: http://localhost:9999/appinfo?vmc_id=0
	router.GET("/vmcdata/sequences", handler.GetVMCSequence)
	router.GET("/topo/show", handler.TopoShow)
	router.POST("/topo/insert", handler.TopoInsert)
	router.POST("/topo/delete", handler.TopoDelete)
	router.POST("/vmc/failure_over", handler.FailureOver)
	router.GET("/vmcdata/get_failure_over_info", handler.GetFailureOverInfo)
	router.GET("/vmcdata/get_failure_over_vmcdata", handler.GetFailureOverVMCData)
	router.POST("/satellite/control/orbit_normal", handler.ApiOrbitNormal)
	router.POST("/satellite/control/orbit_coordinate", handler.ApiOrbitCoordinate)
	router.POST("/satellite/control/marker_coordinate", handler.ApiMarkerCoordinates)
	router.POST("/satellite/control/show_picture", handler.ApiShowPicture)
	router.Run(":9999")

}
