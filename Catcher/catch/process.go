package catch

import (
	"FindGhost/Catcher/conf"
	"FindGhost/Catcher/models"
	"bufio"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpStreamFactory struct{}

type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transPort gopacket.Flow) tcpassembly.Stream {
	hStream := &httpStream{
		net:       net,
		transport: transPort,
		r:         tcpreader.NewReaderStream(),
	}

	go hStream.run()
	return &hStream.r
}

// 处理http数据
func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			return
		} else if err == nil {
			defer req.Body.Close()
			srcIp, destIp := Net2Ips(h.net)
			srcPort, destPort := Transport2Ports(h.transport)

			// 新建http信息模型实例
			httpInfo, _ := models.NewHttpInfo(srcIp, srcPort, destIp, destPort, req)

			// 发送http信息到审计系统
			go func(apiUrl string, httpInfo *models.HttpInfo) {
				if !CheckSelfHttp(apiUrl, httpInfo) {
					logrus.Info("send httpInfo to audit server")
					// 发送httpInfo到审计系统
					_ = SendHttpInfo(httpInfo)
				}
			}(ApiUrl, httpInfo)
		}
	}
}

// 处理抓取到的包
func ProcessPacket(packets chan gopacket.Packet) {
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	ticker := time.Tick(time.Minute)

	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}

			// 处理ip信息
			ProcessIp(packet)

			// 处理dns信息
			ProcessDns(packet)

			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				continue
			}

			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}

// 提取tcp重组流中的ip
func Net2Ips(net gopacket.Flow) (srcIp, destIp string) {
	ips := strings.Split(net.String(), "->")
	if len(ips) > 1 {
		srcIp = ips[0]
		destIp = ips[1]
	}

	return srcIp, destIp
}

// 提取tcp重组流中的port
func Transport2Ports(transport gopacket.Flow) (srcPort, destPort string) {
	ports := strings.Split(transport.String(), "->")
	if len(ports) > 1 {
		srcPort = ports[0]
		destPort = ports[1]
	}
	return srcPort, destPort
}

// 检查http中的host是否是审计系统
func CheckSelfHttp(apiUrl string, req *models.HttpInfo) bool {
	urlParse, err := url.Parse(apiUrl)
	if err == nil {
		apiHost := urlParse.Host
		if apiHost == req.Host {
			return true
		}
	}

	return false
}

// 检查ipInfo中的ip信息
func CheckSelfIp(ApiUrl string, ipInfo *models.IpInfo) bool {
	urlParse, err := url.Parse(ApiUrl)
	if err == nil {
		apiHost := urlParse.Host
		apiIp := strings.Split(apiHost, ":")[0]
		sensorIp := conf.Ips[0]

		if (ipInfo.SrcIp == sensorIp && ipInfo.DestIp == apiIp) || (ipInfo.SrcIp == apiIp && ipInfo.DestIp == sensorIp) {
			return true
		}
	}

	return false
}

// 处理ip信息
func ProcessIp(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		if ip != nil {
			switch ip.Protocol {
			case layers.IPProtocolTCP:
				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcp, _ := tcpLayer.(*layers.TCP)

					srcPort := tcp.SrcPort.String()
					destPort := tcp.DstPort.String()

					ipInfo := models.NewIpInfo("tcp", ip.SrcIP.String(), srcPort, ip.DstIP.String(), destPort)

					// 将ip信息发送到审计系统
					go func(apiUrl string, ipInfo *models.IpInfo) {
						if tcp.SYN && !tcp.ACK && !CheckSelfIp(apiUrl, ipInfo) {
							logrus.Info("send ipInfo to audit server")
							_ = SendIpInfo(ipInfo)
						}
					}(ApiUrl, ipInfo)
				}
			}
		}
	}
}

// 处理dns信息
func ProcessDns(packet gopacket.Packet) {
	var eth layers.Ethernet
	var ip layers.IPv4
	var udp layers.UDP
	var dns layers.DNS
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip, &udp, &dns)
	decodeLayers := make([]gopacket.LayerType, 0)
	err := parser.DecodeLayers(packet.Data(), &decodeLayers)
	if err != nil {
		return
	}

	srcIp := ip.SrcIP
	destIp := ip.DstIP
	for _, q := range dns.Questions {
		dnsInfo := models.NewDnsInfo(q.Type.String(), string(q.Name), srcIp.String(), destIp.String())
		go func(apiUrl string, dnsInfo *models.DnsInfo) {
			logrus.Info("send dnsinfo to audit server")
			_ = SendDnsInfo(dnsInfo)
		}(ApiUrl, dnsInfo)
	}
}
