package models

import (
	"net/http"
	"net/url"
	"time"
)

type HttpReq struct {
	Host          string
	Ip            string
	Client        string
	Port          string
	URL           *url.URL
	Header        http.Header
	RequestURI    string
	Method        string
	ReqParameters url.Values
}

type EvilHttpReq struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	IsEvil   bool      `bson:"is_evil"`
	Data     HttpReq   `bson:"data"`
}
