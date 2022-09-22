package models

import (
	"FindGhost/Analyser/conf"
	"FindGhost/Analyser/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client    *mongo.Client
	Database  *mongo.Database
	DATA_TYPE string
	DATA_HOST string
	DATA_PORT string
	DATA_NAME string
	err       error
)

func init() {
	DATA_TYPE = conf.Cfg.Section("DATABASE").Key("DATA_TYPE").MustString("mongodb")
	DATA_HOST = conf.Cfg.Section("DATABASE").Key("DATA_HOST").MustString("127.0.0.1")
	DATA_PORT = conf.Cfg.Section("DATABASE").Key("DATA_PORT").MustString("27017")
	DATA_NAME = conf.Cfg.Section("DATABASE").Key("DATA_NAME").MustString("test")

	ConnectDB()
}

// 连接数据库
func ConnectDB() {
	uri := DATA_TYPE + "://" + DATA_HOST + ":" + DATA_PORT
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到mongodb
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		util.Log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		util.Log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 连接到数据库
	Database = client.Database(DATA_NAME)

	fmt.Printf("Connected to %v\n", DATA_NAME)
}

// 断开数据库
func DisconnectDB() {
	err = client.Disconnect(context.TODO())
	if err != nil {
		util.Log.Fatal(err)
	}

	fmt.Println("DisConnected MongoDB!")
}
