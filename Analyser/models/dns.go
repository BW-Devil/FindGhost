package models

import (
	"FindGhost/Analyser/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// dns模型
type DnsInfo struct {
	DnsType string `json:"dnsType"`
	DnsName string `json:"dnsName"`
	SrcIp   string `json:"srcIp"`
	DestIp  string `json:"destIp"`
}

type DomainList struct {
	Id     int64    `json:"id"`
	Domain string   `json:"domain"`
	Info   []Source `json:"info"`
}

type DomainListApi struct {
	Evil bool       `json:"evil"`
	Data DomainList `json:"data"`
}

// 存入数据库中的dns模型
type EvilDnsInfo struct {
	Time       time.Time `bson:"time"`
	SensorIp   string    `bson:"sensorIp"`
	DnsType    string    `bson:"dnsType"`
	SrcIp      string    `bson:"srcIp"`
	DestIp     string    `bson:"destIp"`
	EvilDomain string    `bson:"evilDomain"`
	Info       []Source  `bson:"info"`
}

// 创建一个新的恶意dns实例
func NewEvilDnsInfo(sensorIp string, dnsInfo DnsInfo, auditDnsRes DomainListApi) (evilDnsInfo *EvilDnsInfo) {
	time := time.Now()
	return &EvilDnsInfo{Time: time, SensorIp: sensorIp, DnsType: dnsInfo.DnsType,
		SrcIp: dnsInfo.SrcIp, DestIp: dnsInfo.DestIp, EvilDomain: auditDnsRes.Data.Domain,
		Info: auditDnsRes.Data.Info}
}

// 插入恶意dnsInfo到数据库中
func (e *EvilDnsInfo) Insert() {
	isExit := e.Exist()
	if !isExit {
		insertResult, err := Database.Collection("evildomain_list").InsertOne(context.TODO(), e)
		if err != nil {
			util.Log.Fatal(err)
		}
		fmt.Println("success insert new evil domain, it objectId: ", insertResult.InsertedID)
	}
}

// 判断dnsInfo是否已经存在
func (e *EvilDnsInfo) Exist() bool {
	evilDomain := e.EvilDomain
	evilDnsInfos := make([]EvilDnsInfo, 0)

	cur, err := Database.Collection("evildomain_list").Find(context.TODO(), bson.D{{"evilDomain", evilDomain}})
	if err != nil {
		util.Log.Fatal(err)
	}

	// 遍历查询的元素
	for cur.Next(context.TODO()) {
		var evilDnsInfo EvilDnsInfo
		err := cur.Decode(&evilDnsInfo)
		if err != nil {
			util.Log.Fatal(err)
		}

		evilDnsInfos = append(evilDnsInfos, evilDnsInfo)
	}

	if err := cur.Err(); err != nil {
		util.Log.Fatal(err)
	}

	// 关闭游标
	cur.Close(context.TODO())

	if len(evilDnsInfos) > 0 {
		return true
	}

	return false
}

// 列出所有的恶意dnsInfo
func ListEvilDns() []EvilDnsInfo {
	evilDnsInfos := make([]EvilDnsInfo, 0)

	cur, err := Database.Collection("evildomain_list").Find(context.TODO(), bson.D{{}})
	if err != nil {
		util.Log.Fatal(err)
	}

	// 遍历查询的元素
	for cur.Next(context.TODO()) {
		var evilDnsInfo EvilDnsInfo
		err := cur.Decode(&evilDnsInfo)
		if err != nil {
			util.Log.Fatal(err)
		}

		evilDnsInfos = append(evilDnsInfos, evilDnsInfo)
	}

	if err := cur.Err(); err != nil {
		util.Log.Fatal(err)
	}

	// 关闭游标
	cur.Close(context.TODO())

	return evilDnsInfos
}
