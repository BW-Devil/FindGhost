package models

type Dns struct {
	DnsType string `json:"dns_type"`
	DnsName string `json:"dns_name"`
	SrcIp   string `json:"src_ip"`
	DstIp   string `json:"dst_ip"`
}

func NewDns(srcIp, dstIp, dnsType, DnsName string) *Dns {
	return &Dns{SrcIp: srcIp, DstIp: dstIp, DnsType: dnsType, DnsName: DnsName}
}
