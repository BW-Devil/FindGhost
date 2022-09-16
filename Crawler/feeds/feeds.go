package feeds

import (
	"Crawler/gui"
	"Crawler/models"
	"Crawler/util/log"
	"Crawler/web"
	"fmt"
	"github.com/urfave/cli/v2"
	"sync"
	"time"
)

type EvilIpFunc func() (evilIps models.EvilIps, err error)

type EvilDomainFunc func() (evilDomains models.EvilDomains, err error)

var (
	EvilIpFuncs     []EvilIpFunc
	EvilDomainFuncs []EvilDomainFunc
)

func Init() {
	// 获取恶意domain的函数
	EvilDomainFuncs = append(EvilDomainFuncs, FetchDataFromNetLab360)

	// 获取恶意ip的函数
	EvilIpFuncs = append(EvilIpFuncs, FetchFromGreensnow)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromRutgers)
}

// 获取恶意domain
func FetchEvilDomains() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(len(EvilDomainFuncs))
	for _, fn := range EvilDomainFuncs {
		go func(fn EvilDomainFunc) {
			models.SaveEvilDomains(fn())
			wg.Done()
		}(fn)
	}
	wg.Wait()
	log.Log.Infof("Fetch evil domain finish, speed time : %v\n", time.Since(startTime))
}

// 获取恶意ip
func FetchEvilIps() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(len(EvilIpFuncs))
	fmt.Println(len(EvilIpFuncs))
	for _, fn := range EvilIpFuncs {
		go func(fn EvilIpFunc) {
			models.SaveEvilIps(fn())
			wg.Done()
		}(fn)
	}
	wg.Wait()
	log.Log.Infof("Fetch evil ip finish, speed time : %v\n", time.Since(startTime))
}

// 爬取所有数据
func FetchAll(cCtx *cli.Context) {
	for {
		go func(ctx *cli.Context) {
			FetchEvilIps()
			FetchEvilDomains()
			models.Status()
			//models.SaveFile(ctx)
		}(cCtx)

		time.Sleep(60 * 60 * time.Second)
	}
}

// 启动web程序
func StartUpWeb(cCtx *cli.Context) (err error) {
	Init()
	go FetchAll(cCtx)
	web.RunWeb(cCtx)
	return err
}

// 启动gui程序
func StartUpGui(cCtx *cli.Context) (err error) {
	gui.StartUp(cCtx)
	return err
}

// 导出缓冲区中的数据到本地
func Dump(cCtx *cli.Context) (err error) {
	Init()
	FetchEvilDomains()
	models.Status()
	models.SaveFile(cCtx)
	return err
}
