package models

// ip模型
type IpInfo struct {
	Protocol string `json:"protocol"`
	SrcIp    string `json:"srcIp"`
	SrcPort  string `json:"srcPort"`
	DestIp   string `json:"destIp"`
	DestPort string `json:"destPort"`
}

// 创建新的ip模型
func NewIpInfo(protocal, srcIp, srcPort, destIp, destPort string) (ipInfo *IpInfo) {
	return &IpInfo{Protocol: protocal, SrcIp: srcIp, SrcPort: srcPort, DestIp: destIp, DestPort: destPort}
}
