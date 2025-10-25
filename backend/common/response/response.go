// Package response 响应模型
//
// @File: response.go
// @Desc: 响应模型
// @Date: 2025-10-25

package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response 统一接口返回参数
func Response(w http.ResponseWriter, data interface{}, err error) {
	var resp Resp
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Msg = "OK"
		resp.Data = data
	}
	httpx.OkJson(w, resp)
}
