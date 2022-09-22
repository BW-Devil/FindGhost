package models

import (
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
