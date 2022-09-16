package gui

import (
	"Crawler/web/router"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/urfave/cli/v2"
	"image/color"
	"io"
	"net"
	"net/http"
)

var (
	domainDetail      router.DomainListApi
	ipDetail          router.IplistApi
	resIpOrDomain     string
	resSrc            string = ""
	resIpOrDomainText *canvas.Text
	resSrcText        *canvas.Text
)

// 启动gui
func StartUp(cCtx *cli.Context) {
	a := app.New()

	// 新建窗口
	w := a.NewWindow("FindGhost")

	// 标题
	title := canvas.NewText("FindGhost Crawler", color.White)
	title.Alignment = fyne.TextAlignCenter

	// 查找框
	entry := widget.NewEntry()
	entry.SetPlaceHolder("please input ip or domain")

	entry.Wrapping = fyne.TextWrapOff
	entry.ActionItem = widget.NewButton("Find", func() {
		CheckInput(entry.Text)
		resIpOrDomain = ""
		resSrc = ""
		// 查询结果
		if domainDetail.Evil {
			resIpOrDomain = fmt.Sprintf("Evil Domain: %v", domainDetail.Data.Domain)
			resSrc = resSrc + domainDetail.Data.Info[0].Source
			resSrc = fmt.Sprintf("Data From: %v", resSrc)
		} else if ipDetail.Evil {
			resIpOrDomain = fmt.Sprintf("Evil Ip: %v", ipDetail.Data.Ip)
			resSrc = resSrc + ipDetail.Data.Info[0].Source
			resSrc = fmt.Sprintf("Data From: %v", resSrc)
		} else {
			resIpOrDomain = fmt.Sprintf("It's not evil")
			resSrc = fmt.Sprintf("It's not evil")
		}

		resIpOrDomainText = canvas.NewText(resIpOrDomain, nil)
		resSrcText = canvas.NewText(resSrc, nil)

		resContent := container.NewGridWithRows(2, resIpOrDomainText, resSrcText)

		// 结果弹窗
		resDialog := dialog.NewCustom("Find Result", "dismiss", resContent, w)
		resDialog.SetDismissText("close")
		resDialog.Show()

		domainDetail = router.DomainListApi{}
		ipDetail = router.IplistApi{}
		entry.Text = ""

	})

	// 页面布局
	//content := container.New(layout.NewCenterLayout(), title, entry)
	finalContent := container.NewVBox(title, entry)

	// 设置窗口内容
	w.SetContent(finalContent)

	// 设置窗口大小
	w.Resize(fyne.NewSize(1000, 500))
	w.SetFixedSize(true)

	w.ShowAndRun()
}

// 验证是否为恶意ip或恶意domain
func CheckInput(s string) {
	// 判断传入的是Ip还是domain
	isIp := net.ParseIP(s)

	if isIp == nil {
		// 查询输入的值是否是恶意domain
		domainUrl := fmt.Sprintf("http://127.0.0.1:8888/api/domain/%v", s)
		domainResp, err := http.Get(domainUrl)
		if err == nil {
			res, err := io.ReadAll(domainResp.Body)
			if err == nil {
				err = json.Unmarshal(res, &domainDetail)
			}
		}
	} else {
		// 查询输入的值是否是恶意ip
		ipUrl := fmt.Sprintf("http://127.0.0.1:8888/api/ip/%v", s)
		ipResp, err := http.Get(ipUrl)
		if err == nil {
			res, err := io.ReadAll(ipResp.Body)
			if err == nil {
				err = json.Unmarshal(res, &ipDetail)
			}
		}
	}

	fmt.Println(domainDetail)
	fmt.Println(ipDetail)
}
