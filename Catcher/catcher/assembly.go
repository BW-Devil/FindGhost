package catcher

import (
	"FindGhost/Catcher/models"
	"FindGhost/Catcher/util"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
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

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hStream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hStream.run()
	return &hStream.r
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			return
		} else if err == nil {
			defer req.Body.Close()

			clientIp, dstIp := SplitNet2Ips(h.net)
			srcPort, dstPort := Transport2Ports(h.transport)

			httpReq, _ := models.NewHttpReq(req, clientIp, dstIp, dstPort)
			// send to sever
			go func(u string, req *models.HttpReq) {
				if !CheckSelfHtml(u, req) {
					util.Log.Warnf("%v:%v -> %v(%v:%v), %v, %v, %v, req_param: %v, req_body: %v", httpReq.Client, srcPort, httpReq.Host, httpReq.Ip,
						httpReq.Port, httpReq.Method, httpReq.URL, httpReq.Header, httpReq.ReqParameters, httpReq.RequestBody)
					_ = SendHTML(req)
				}
			}(ApiUrl, httpReq)
		}
	}
}

func SplitNet2Ips(net gopacket.Flow) (client, host string) {
	ips := strings.Split(net.String(), "->")
	if len(ips) > 1 {
		client = ips[0]
		host = ips[1]
	}
	return client, host
}

func Transport2Ports(transport gopacket.Flow) (src, dst string) {
	ports := strings.Split(transport.String(), "->")
	if len(ports) > 1 {
		src = ports[0]
		dst = ports[1]
	}
	return src, dst
}

func CheckSelfHtml(ApiUrl string, req *models.HttpReq) (ret bool) {
	urlParsed, err := url.Parse(ApiUrl)
	if err == nil {
		apiIp := urlParsed.Host
		if apiIp == req.Host {
			ret = true
		}
		// util.Log.Errorf("apiIp: %v, req.Host: %v, ret: %v", apiIp, req.Host, ret)
	}
	return ret
}

func ProcessPackets(packets chan gopacket.Packet) {
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

			// 处理tcp/udp数据包
			processPacket(packet)
			// 处理DNS包
			parseDNS(packet)

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

func SendHTML(req *models.HttpReq) error {
	reqJson, err := json.Marshal(req)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	urlApi := fmt.Sprintf("%v%v", ApiUrl, "/api/http/")
	secureKey := util.MakeSign(timestamp, SecureKey)
	_, err = http.PostForm(urlApi, url.Values{"timestamp": {timestamp}, "secureKey": {secureKey}, "data": {string(reqJson)}})
	return err
}

func SendDns(dns *models.Dns) error {
	reqJson, err := json.Marshal(dns)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	urlApi := fmt.Sprintf("%v%v", ApiUrl, "/api/dns/")
	secureKey := util.MakeSign(timestamp, SecureKey)
	_, err = http.PostForm(urlApi, url.Values{"timestamp": {timestamp}, "secureKey": {secureKey}, "data": {string(reqJson)}})
	return err
}
