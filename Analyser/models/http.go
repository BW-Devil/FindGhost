package models

import (
	"FindGhost/Analyser/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/url"
	"time"
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

// 恶意httpInfo模型
type EvilHttpInfo struct {
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensorIp"`
	EvilIp   string    `bson:"evilIp"`
	Info     []Source  `bson:"info"`
	Data     HttpInfo  `bson:"data"`
}

// 创建EvilHttpInfo实例
func NewEvilHttpInfo(sensorIp string, httpInfo HttpInfo, auditIpRes IpListApi) (evilHttpInfo *EvilHttpInfo) {
	time := time.Now()
	return &EvilHttpInfo{Time: time, SensorIp: sensorIp, EvilIp: auditIpRes.Data.Ip,
		Info: auditIpRes.Data.Info, Data: httpInfo}
}

// 将恶意http实例插入数据库
func (e *EvilHttpInfo) Insert() {
	isExit := e.Exist()
	if !isExit {
		insertResult, err := Database.Collection("evilhttp_list").InsertOne(context.TODO(), e)
		if err != nil {
			util.Log.Fatal(err)
		}
		fmt.Println("success insert new evil http, it objectId: ", insertResult.InsertedID)
	}
}

// 判断http是否已经存在
func (e *EvilHttpInfo) Exist() bool {
	evilIp := e.EvilIp
	evilHttpInfos := make([]EvilHttpInfo, 0)

	cur, err := Database.Collection("evilhttp_list").Find(context.TODO(), bson.D{{"evilIp", evilIp}})
	if err != nil {
		util.Log.Fatal(err)
	}

	// 遍历查询的元素
	for cur.Next(context.TODO()) {
		var evilHttpInfo EvilHttpInfo
		err := cur.Decode(&evilHttpInfo)
		if err != nil {
			util.Log.Fatal(err)
		}

		evilHttpInfos = append(evilHttpInfos, evilHttpInfo)
	}

	if err := cur.Err(); err != nil {
		util.Log.Fatal(err)
	}

	// 关闭游标
	cur.Close(context.TODO())

	if len(evilHttpInfos) > 0 {
		return true
	}

	return false
}

// 列出所有evilHttpInfo
func ListEvilHttp() []EvilHttpInfo {
	evilHttpInfos := make([]EvilHttpInfo, 0)

	cur, err := Database.Collection("evilhttp_list").Find(context.TODO(), bson.D{{}})
	if err != nil {
		util.Log.Fatal(err)
	}

	// 遍历查询的元素
	for cur.Next(context.TODO()) {
		var evilHttpInfo EvilHttpInfo
		err := cur.Decode(&evilHttpInfo)
		if err != nil {
			util.Log.Fatal(err)
		}

		evilHttpInfos = append(evilHttpInfos, evilHttpInfo)
	}

	if err := cur.Err(); err != nil {
		util.Log.Fatal(err)
	}

	// 关闭游标
	cur.Close(context.TODO())

	return evilHttpInfos
}
