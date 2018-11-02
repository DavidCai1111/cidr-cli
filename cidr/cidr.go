package cidr

import (
	"fmt"
	"net"
)

// IPAddressInfo - IP adress infomation
type IPAddressInfo struct {
	IPAvailable      int    `json:"IPAvailable"`
	NetID            string `json:"netID"`
	Mask             string `json:"mask"`
	FirstIPAvailable string `json:"firstIPAvailable"`
	LastIPAvailable  string `json:"lastIPAvailable"`
	Boardcast        string `json:"boardcast"`
}

// Parse - Parse the given CIDR to IP adress infomation
func Parse(cidr string) (*IPAddressInfo, error) {
	ipv4Addr, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	info := &IPAddressInfo{
		NetID:            ipv4Net.IP.String(),
		Mask:             ipv4Net.Mask.String(),
		FirstIPAvailable: NextIP(&ipv4Net.IP).String(),
	}

	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	fmt.Printf("%+v/n", info)

	return info, nil
}

// NextIP - get next IP address
func NextIP(ip *net.IP) *net.IP {
	newIP := ip.To4()
	newIP[3]++

	return &newIP
}
