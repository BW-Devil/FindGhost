package models

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpReq struct {
	Host          string
	Ip            string
	Client        string
	Port          string
	URL           *url.URL
	Header        http.Header
	RequestURI    string
	RequestBody   string
	Method        string
	ReqParameters url.Values
}

func NewHttpReq(req *http.Request, client string, ip string, port string) (httpReq *HttpReq, err error) {
	err = req.ParseForm()
	body := req.Body
	buff, err := ioutil.ReadAll(body)
	return &HttpReq{Host: req.Host, Client: client, Ip: ip, Port: port, URL: req.URL, Header: req.Header,
		RequestURI: req.RequestURI, RequestBody: string(buff), Method: req.Method, ReqParameters: req.Form}, err
}
