package models

import "time"

type Dns struct {
	DnsType string `json:"dns_type"`
	DnsName string `json:"dns_name"`
	SrcIp   string `json:"src_ip"`
	DstIp   string `json:"dst_ip"`
}

type EvilDns struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	IsEvil   bool      `bson:"is_evil"`
	Data     Dns       `bson:"data"`
}
