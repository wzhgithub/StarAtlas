package model

import "fmt"

type CommonResponse struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCommonResponseFail(err error) *CommonResponse {
	if err == nil {
		err = fmt.Errorf("unknown error")
	}

	return &CommonResponse{
		Code: 500,
		Msg:  err.Error(),
		Data: nil,
	}
}

func NewCommonResponseSucc(data interface{}) *CommonResponse {
	return &CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}
