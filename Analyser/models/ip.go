package models

import (
	"FindGhost/Analyser/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ipInfo模型
type IpInfo struct {
	Protocol string `json:"protocol"`
	SrcIp    string `json:"srcIp"`
	SrcPort  string `json:"srcPort"`
	DestIp   string `json:"destIp"`
	DestPort string `json:"destPort"`
}

// 恶意ip来源信息
type Source struct {
	Desc   string `json:"desc"`
	Source string `json:"source"`
}

type IpList struct {
	Id   int64    `json:"id"`
	Ip   string   `json:"ip"`
	Info []Source `json:"info"`
}

type IpListApi struct {
	Evil bool   `json:"evil"`
	Data IpList `json:"data"`
}

// 存入数据库的恶意ip模型
type EvilIpInfo struct {
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensorIp"`
	EvilIp   string    `bson:"evilIp"`
	Info     []Source  `bson:"info"`
	Protocol string    `bson:"protocol"`
	SrcIp    string    `bson:"srcIp"`
	SrcPort  string    `bson:"srcPort"`
	DestIp   string    `bson:"destIp"`
	DestPort string    `bson:"destPort"`
}

// 创建一个新的恶意ip实例
func NewEvilIpInfo(sensorIp string, ipInfo IpInfo, auditIpRes IpListApi) (evilIpInfo *EvilIpInfo) {
	time := time.Now()
	return &EvilIpInfo{Time: time, SensorIp: sensorIp, EvilIp: auditIpRes.Data.Ip,
		Info: auditIpRes.Data.Info, Protocol: ipInfo.Protocol, SrcIp: ipInfo.SrcIp,
		SrcPort: ipInfo.SrcPort, DestIp: ipInfo.DestIp, DestPort: ipInfo.DestPort}
}

// 将恶意ip实例插入数据库
func (e *EvilIpInfo) Insert() {
	isExit := e.Exist()
	if !isExit {
		insertResult, err := Database.Collection("evilip_list").InsertOne(context.TODO(), e)
		if err != nil {
			util.Log.Fatal(err)
		}
		fmt.Println("success insert new evil ip, it objectId: ", insertResult.InsertedID)
	}
}

// 判断ip是否已经存在
func (e *EvilIpInfo) Exist() bool {
	evilIp := e.EvilIp
	evilIpInfos := make([]EvilIpInfo, 0)

	cur, err := Database.Collection("evilip_list").Find(context.TODO(), bson.D{{"evilIp", evilIp}})
	if err != nil {
		util.Log.Fatal(err)
	}

	// 遍历查询的元素
	for cur.Next(context.TODO()) {
		var evilIpInfo EvilIpInfo
		err := cur.Decode(&evilIpInfo)
		if err != nil {
			util.Log.Fatal(err)
		}

		evilIpInfos = append(evilIpInfos, evilIpInfo)
	}

	if err := cur.Err(); err != nil {
		util.Log.Fatal(err)
	}

	// 关闭游标
	cur.Close(context.TODO())

	if len(evilIpInfos) > 0 {
		return true
	}

	return false
}
