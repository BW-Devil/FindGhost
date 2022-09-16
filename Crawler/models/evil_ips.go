package models

import (
	"Crawler/util/log"
	"context"
	"fmt"
)

// 恶意ip/domain来源信息
type Source struct {
	Desc   string
	Source string
}

// 恶意ip
type EvilIps struct {
	Ips []string
	Src Source
}

// 恶意domains
type EvilDomains struct {
	Domains []string
	Src     Source
}

type IpList struct {
	Id   int64
	Ip   string   `bson:"ip"`
	Info []Source `bson:"info"`
}

type DomainList struct {
	Id     int64
	Domain string   `bson:"domain"`
	Info   []Source `bson:"info"`
}

func NewIpList(ip string, info []Source) IpList {
	infos := make([]Source, 0)
	infos = append(infos, info...)
	return IpList{Ip: ip, Info: infos}
}

func NewDomainList(domain string, info []Source) DomainList {
	infos := make([]Source, 0)
	infos = append(infos, info...)
	return DomainList{Domain: domain, Info: infos}
}

func InsertIps2DB(ips []IpList) {
	for _, ip := range ips {
		insertResult, err := Database.Collection("ip_list").InsertOne(context.TODO(), ip)
		if err != nil {
			log.Log.Fatal(err)
		}

		fmt.Println("success insert ip: ", insertResult.InsertedID)
	}
}

func InsertDomains2DB(domains []DomainList) {
	for _, domain := range domains {
		insertResult, err := Database.Collection("domain_list").InsertOne(context.TODO(), domain)
		if err != nil {
			log.Log.Fatal(err)
		}

		fmt.Println("success insert domain: ", insertResult.InsertedID)
	}
}
