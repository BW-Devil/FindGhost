package catcher

import (
	"FindGhost/Catcher/conf"
	"FindGhost/Catcher/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"time"
)

var (
	device      string
	snapshotLen int32         = 1024
	promiscuous bool          = true
	timeout     time.Duration = pcap.BlockForever
	handle      *pcap.Handle
	err         error

	DebugMode bool
	filter    = ""

	ApiUrl    string
	SecureKey string
)

func init() {
	device = conf.DeviceName
	DebugMode = conf.DebugMode
	filter = conf.FilterRule

	ApiUrl = conf.Cfg.Section("SEVER").Key("API_URL").MustString("")
	SecureKey = conf.Cfg.Section("SEVER").Key("API_KEY").MustString("")

}

func Start(ctx *cli.Context) error {
	if ctx.IsSet("debug") {
		DebugMode = ctx.Bool("debug")
	}
	if DebugMode {
		util.Log.Logger.Level = logrus.DebugLevel
	}

	if ctx.IsSet("length") {
		snapshotLen = int32(ctx.Int("len"))
	}

	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		util.Log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	if ctx.IsSet("filter") {
		filter = ctx.String("filter")
	}
	err := handle.SetBPFFilter(filter)
	util.Log.Infof("set SetBPFFilter: %v, err: %v", filter, err)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	ProcessPackets(packetSource.Packets())
	return nil
}
