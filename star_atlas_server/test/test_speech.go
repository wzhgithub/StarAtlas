package main

import (
	"fmt"
	"star_atlas_server/config"
	"star_atlas_server/handler"

	"github.com/golang/glog"
	"github.com/nl8590687/asrt-sdk-go/sdk"
)

func main() {
	err := config.Init("../config.yaml")
	glog.Infof("config:%+v\n", config.CommonConfig)
	if err != nil {
		glog.Fatalf(err.Error())
	}
	filename := "record.wav"
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
