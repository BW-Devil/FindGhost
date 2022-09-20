package models

import (
	"io"
	"net/http"
	"net/url"
)

// http信息模型
type HttpInfo struct {
	SrcIp         string      `json:"srcIp"`
	SrcPort       string      `json:"srcPort"`
	DestIp        string      `json:"destIp"`
	DestPort      string      `json:"destPort"`
	Host          string      `json:"host"`
	Method        string      `json:"method"`
	RequestURI    string      `json:"requestURI"`
	Header        http.Header `json:"header"`
	RequestBody   string      `json:"requestBody"`
	ReqParameters url.Values  `json:"reqParameters"`
}

// 新建http信息模型
func NewHttpInfo(srcIp, srcPort, destIp, destPort string, req *http.Request) (httpInfo *HttpInfo, err error) {
	err = req.ParseForm()
	requestBody, err := io.ReadAll(req.Body)
	return &HttpInfo{SrcIp: srcIp, SrcPort: srcPort, DestIp: destIp, DestPort: destPort, Host: req.Host,
		Method: req.Method, RequestURI: req.RequestURI, Header: req.Header, RequestBody: string(requestBody),
		ReqParameters: req.Form}, err
}
