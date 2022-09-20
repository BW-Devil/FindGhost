package models

// dns模型
type DnsInfo struct {
	DnsType string `json:"dnsType"`
	DnsName string `json:"dnsName"`
	SrcIp   string `json:"srcIp"`
	DestIp  string `json:"destIp"`
}

// 创建新的dns模型
func NewDnsInfo(dnsType, dnsName, srcIp, destIp string) (dnsInfo *DnsInfo) {
	return &DnsInfo{DnsType: dnsType, DnsName: dnsName, SrcIp: srcIp, DestIp: destIp}
}
