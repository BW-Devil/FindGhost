package util

import (
	"crypto/md5"
	"fmt"
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
