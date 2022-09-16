package feeds

import (
	"Crawler/models"
	"Crawler/util/client"
	"io"
	"strings"
)

func FetchFromRutgers() (evilIps models.EvilIps, err error) {
	url := "https://report.cs.rutgers.edu/DROP/attackers"
	src := "rutgers.edu"
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
