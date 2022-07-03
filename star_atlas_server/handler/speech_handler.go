package handler

import (
	"encoding/json"
	"fmt"
	"star_atlas_server/config"

	"github.com/golang/glog"
	"github.com/nl8590687/asrt-sdk-go/common"
)

const cContenttype = "application/json"
const cPostMethod = "POST"
const cWavDataMaxLength = 16000 * 2 * 16

func recogniteSpeech(wavData []byte, frameRate int, channels int, byteWidth int) (*common.AsrtAPIResponse, error) {
	if len(wavData) > cWavDataMaxLength {
		return nil, fmt.Errorf("error: %s `%d`, %s `%d`",
			"Too long wave sample byte length:", len(wavData),
			"the max length is", cWavDataMaxLength)
	}

	requestBody := common.AsrtAPISpeechRequest{
		Samples:    common.BytesToBase64(wavData),
		SampleRate: frameRate,
		Channels:   channels,
		ByteWidth:  byteWidth,
	}

	byteForm, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/speech", config.CommonConfig.SpeechURL)
	rspBody, err := common.SendHTTPRequest(url, cPostMethod, byteForm, cContenttype)
	if err != nil {
		return nil, err
	}

	responseBody := common.AsrtAPIResponse{}
	err = json.Unmarshal(rspBody, &responseBody)
	if err != nil {
		return nil, err
	}

	glog.Infof("recv speech result: %+v", responseBody)
	return &responseBody, nil
}
