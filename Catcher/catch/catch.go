package catch

import (
	"FindGhost/Catcher/conf"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/urfave/cli/v2"
	"log"
	"time"
)

var (
	device      string        = "eth0"
	snapshotLen int32         = 1024
	promiscuous bool          = true
	timeout     time.Duration = pcap.BlockForever
	err         error
	handle      *pcap.Handle

	DeviceName string
	DebugMode  bool
	FilterRule string = ""
	ApiUrl     string
	ApiKey     string
)

func init() {
	cfg := conf.Cfg
	DeviceName = conf.DeviceName
	DebugMode = conf.DebugMode
	FilterRule = conf.FilterRule

	ApiUrl = cfg.Section("SERVER").Key("API_URL").MustString("")
	ApiKey = cfg.Section("SERVER").Key("API_KEY").MustString("")
}

func Start(cCtx *cli.Context) error {
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤规则
	handle.SetBPFFilter(FilterRule)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	ProcessPacket(packetSource.Packets())
	return nil
}
