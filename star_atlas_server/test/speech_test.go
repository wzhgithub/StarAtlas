package test

import (
	"star_atlas_server/handler"
	"testing"

	"github.com/nl8590687/asrt-sdk-go/sdk"
)

func TestSpeechModule(t *testing.T) {
	filename := "record.wav"
	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		t.Log(err)
	}
	res, err := handler.RecogniteByType(wave.GetRawSamples(), 16000, 1, 2, "speech")
	if err != nil {
		t.Log(err)
	}
	t.Logf("received speech from server %v\n", res)
	if res.StatusCode != 200 {
		t.Log(res.StatucMesaage)
	}
	t.Logf("received speech word %v\n", res.Result)
}
