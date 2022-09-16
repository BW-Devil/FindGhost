package models

import (
	client2 "Crawler/util/client"
	"Crawler/util/log"
	"encoding/gob"
	"github.com/patrickmn/go-cache"
	"github.com/urfave/cli/v2"
)

func init() {
	gob.Register(DomainList{})
	gob.Register(IpList{})
}

// 查看指定缓冲区状态
func CacheStatus(cache *cache.Cache) (count int, items map[string]cache.Item) {
	count = cache.ItemCount()
	items = cache.Items()
	return count, items
}

// 查看缓冲区状态
func Status() {
	{
		count, _ := CacheStatus(CACHE_IPS)
		log.Log.Infof("Evil Ips count : %v", count)
	}
	{
		count, _ := CacheStatus(CACHE_DOMAINS)
		log.Log.Infof("Evil Domains count : %v", count)
	}
}

// 将缓冲区数据保存到本地文件
func SaveFile(cCtx *cli.Context) (err error) {
	CACHE_IPS.SaveFile("evil_ips")
	CACHE_DOMAINS.SaveFile("evil_domains")
	return err
}

// 将缓冲区数据存入数据库
func Save2DB() {
	// 连接数据库
	ConnectDB()

	// 清空数据表
	ClearDB()

	// 断开数据库
	defer DisconnectDB()

	Num := 500

	// 将缓冲区domain存入数据库
	domainList := make([]DomainList, 0)
	{
		count, items := CacheStatus(CACHE_DOMAINS)
		for _, v := range items {
			d := v.Object.(DomainList)
			domainList = append(domainList, d)
		}

		if count%Num == 0 {
			batch := count / Num
			for i := 0; i < batch; i++ {
				domains := domainList[i*Num : (i+1)*Num]
				InsertDomains2DB(domains)
			}
		} else {
			batch := count / Num
			for i := 0; i < batch; i++ {
				domains := domainList[i*Num : (i+1)*Num]
				InsertDomains2DB(domains)
			}
			InsertDomains2DB(domainList[batch*Num : count])
		}
	}

	// 将缓冲区ip存入数据库
	ipList := make([]IpList, 0)
	{
		count, items := CacheStatus(CACHE_IPS)
		for _, v := range items {
			i := v.Object.(IpList)
			ipList = append(ipList, i)
		}

		if count%Num == 0 {
			batch := count / Num
			for i := 0; i < batch; i++ {
				ips := ipList[i*Num : (i+1)*Num]
				InsertIps2DB(ips)
			}
		} else {
			batch := count / Num
			for i := 0; i < batch; i++ {
				ips := ipList[i*Num : (i+1)*Num]
				InsertIps2DB(ips)
			}
			InsertIps2DB(ipList[batch*Num : count])
		}
	}
}

// 将爬取的domain存入缓冲区
func SaveEvilDomains(evilDomains EvilDomains, err error) {
	if err == nil {
		domains := evilDomains.Domains
		src := evilDomains.Src

		for _, d := range domains {
			infos := make([]Source, 0)
			infos = append(infos, src)
			domain := NewDomainList(d, infos)
			item, found := CACHE_DOMAINS.Get(d)
			if found {
				v := item.(DomainList)
				infos := v.Info

				sliceString := make([]string, 0)
				for _, s := range infos {
					sliceString = append(sliceString, s.Source)
				}

				if !client2.ContainsString(sliceString, src.Source) {
					infos = append(infos, src)
				}

				domain = NewDomainList(d, infos)
			}

			CACHE_DOMAINS.Set(d, domain, cache.NoExpiration)

		}
	}
}

// 将获得的ip存入缓冲区
func SaveEvilIps(evilIps EvilIps, err error) {
	if err == nil {
		ips := evilIps.Ips
		src := evilIps.Src

		for _, ip := range ips {
			infos := make([]Source, 0)
			infos = append(infos, src)
			ipList := NewIpList(ip, infos)

			item, found := CACHE_IPS.Get(ip)
			if found {
				v := item.(IpList)
				infos := v.Info

				sliceString := make([]string, 0)
				for _, s := range infos {
					sliceString = append(sliceString, s.Source)
				}

				if !client2.ContainsString(sliceString, src.Source) {
					infos = append(infos, src)
				}

				ipList = NewIpList(ip, infos)
			}

			CACHE_IPS.Set(ip, ipList, cache.NoExpiration)
		}
	}
}
