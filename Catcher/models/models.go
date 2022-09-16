package models

type ConnectionInfo struct {
	Protocol string `json:"protocol"`
	SrcIp    string `json:"src_ip"`
	SrcPort  string `json:"src_port"`
	DstIp    string `json:"dst_ip"`
	DstPort  string `json:"dst_port"`
}

func NewConnectionInfo(proto string, srcIp string, srcPort string, dstIp string, dstPort string) (connInfo *ConnectionInfo) {
	return &ConnectionInfo{Protocol: proto, SrcIp: srcIp, SrcPort: srcPort, DstIp: dstIp, DstPort: dstPort}
}
