package catch

import (
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

// 处理ip信息
func ProcessIp(packet gopacket.Packet) {
	
}

// 处理dns信息
func ProcessDns(packet gopacket.Packet) {

}