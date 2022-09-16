package feeds

import (
	"Crawler/models"
	"Crawler/util/client"
	"io"
	"strings"
)

func FetchDataFromNetLab360() (evilDomains models.EvilDomains, err error) {
	url := "https://data.netlab.360.com/feeds/dga/dga.txt"
	src := "data.netlab.360.com"
	desc := "360 netlab DGA Domain List"
	check := "\t"

	evilDomains.Src.Source = src
	evilDomains.Src.Desc = desc

	resp, err := client.GetPage(url)
	if err == nil {
		ret, err := io.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.Contains(line, "#") {
					continue
				}
				tmp := strings.Split(line, check)

				if len(tmp) > 1 {
					evilDomains.Domains = append(evilDomains.Domains, tmp[1])
				}
			}
		}
	}

	return evilDomains, err
}
