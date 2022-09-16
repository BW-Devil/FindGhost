package feeds

import (
	"Crawler/models"
	"Crawler/util/client"
	"io"
	"strings"
)

func FetchFromGreensnow() (evilIps models.EvilIps, err error) {
	url := "http://blocklist.greensnow.co/greensnow.txt"
	src := "blocklist.greensnow.co"
	desc := "known attacker"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	resp, err := client.GetPage(url)
	if err == nil {
		ret, err := io.ReadAll(resp)
		if err == nil {
			ips := strings.Split(string(ret), "\n")
			evilIps.Ips = append(evilIps.Ips, ips...)
		}
	}
	return evilIps, err
}
