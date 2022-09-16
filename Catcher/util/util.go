package util

import "github.com/google/gopacket/pcap"

func GetIpList(deviceName string) (ips []string, err error) {
	devices, err := pcap.FindAllDevs()
	if err == nil {
		for _, device := range devices {
			if device.Name == deviceName {
				for _, addr := range device.Addresses {
					if addr.IP.To4() != nil {
						ips = append(ips, addr.IP.To4().String())
					}
				}

			}
		}
	}

	return ips, err
}
