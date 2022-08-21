package test

import (
	"star_atlas_server/config"
	"star_atlas_server/handler"
	"testing"

	"github.com/nl8590687/asrt-sdk-go/sdk"
)

func TestSpeech(t *testing.T) {
	err := config.Init("../config.yaml")
	t.Logf("config:%+v\n", config.CommonConfig)
	if err != nil {
		t.Error(err)
	}
	filename := "record.wav"
	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		t.Error(err)
	}
	t.Logf("byteData: %+v\n", byteData)
	t.Logf("wave.GetRawSamples(): %+v\n", wave.GetRawSamples())
	res, err := handler.RecogniteByType(wave.GetRawSamples(), 16000, 1, 2, "speech")
	if err != nil {
		t.Error(err)
	}
	t.Logf("received speech from server %v\n", res)
	if res.StatusCode != 200 {
		t.Logf(res.StatucMesaage)
	}
	t.Logf("received speech word %v\n", res.Result)
}
