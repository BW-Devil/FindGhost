package models

// dns模型
type DnsInfo struct {
	DnsType string `json:"dnsType"`
	DnsName string `json:"dnsName"`
	SrcIp   string `json:"srcIp"`
	DestIp  string `json:"destIp"`
}
