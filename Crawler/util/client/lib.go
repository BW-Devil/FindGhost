package client

import (
	"Crawler/conf"
	"Crawler/util/log"
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func MD5(s string) (m string) {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MakeSign(t, key string) (sign string) {
	sign = MD5(fmt.Sprintf("%s%s", t, key))
	return sign
}

// 根据url获取页面信息
func GetPage(url string) (io.Reader, error) {
	if conf.DEBUG {
		log.Log.Logger.Infof("Get data from %v", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// 判断string数组中是否包含某个值
func ContainsString(s1 []string, key string) bool {
	for _, v := range s1 {
		if v == key {
			return true
		}
	}

	return false
}
