package models

import (
	"time"
)

type ConnectionInfo struct {
	Protocol string `json:"protocol"`
	SrcIp    string `json:"src_ip"`
	SrcPort  string `json:"src_port"`
	DstIp    string `json:"dst_ip"`
	DstPort  string `json:"dst_port"`
}

// bad ip or dns source info
type Source struct {
	Desc   string `json:"desc"`
	Source string `json:"source"`
}

// evil ips
type EvilIps struct {
	Ips []string `json:"ips"`
	Src Source   `json:"src"`
}

type IpList struct {
	Id   int64
	Ip   string   `json:"ip"`
	Info []Source `json:"info"`
}

type IplistApi struct {
	Evil bool   `json:"evil"`
	Data IpList `json:"data"`
}

type EvilConnectInfo struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	Protocol string    `bson:"protocol"`
	SrcIp    string    `bson:"src_ip"`
	SrcPort  string    `bson:"src_port"`
	DstIp    string    `bson:"dst_ip" `
	DstPort  string    `bson:"dst_port" `
	IsEvil   bool      `bson:"is_evil" `
	Data     []Source  `bson:"data"`
}
