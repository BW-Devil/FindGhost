package util

import (
	"crypto/md5"
	"fmt"
	"github.com/google/gopacket/pcap"
	"io"
)

// md5加密
func MD5(s string) string {
	m := md5.New()
	io.WriteString(m, s)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// 做签名
func MakeSign(timeStamp, key string) string {
	sign := MD5(fmt.Sprintf("%s%s", timeStamp, key))
	return sign
}

// 获取网络设备ip
func GetIps(deviceName string) (ips []string, err error) {
	devices, err := pcap.FindAllDevs()
	if err == nil {
		for _, device := range devices {
			if device.Name == deviceName {
				for _, address := range device.Addresses {
					if address.IP.To4() != nil {
						ips = append(ips, address.IP.To4().String())
					}
				}
			}
		}
	}

	return ips, err
}
