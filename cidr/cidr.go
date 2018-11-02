package cidr

import (
	"math"
	"net"
)

// IPAddressInfo - IP adress information
type IPAddressInfo struct {
	IPAvailable      int    `json:"IPAvailable"`
	NetID            string `json:"netID"`
	Mask             string `json:"mask"`
	FirstIPAvailable string `json:"-"`
	LastIPAvailable  string `json:"-"`
	Boardcast        string `json:"-"`
}

// Parse - Parse the given CIDR to IP adress information
func Parse(cidr string) (*IPAddressInfo, error) {
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	size, bits := ipv4Net.Mask.Size()
	IPAvailable := int(math.Pow(2, float64(bits-size)) - 2)

	info := &IPAddressInfo{
		IPAvailable:      IPAvailable,
		NetID:            ipv4Net.IP.String(),
		Mask:             ipv4Net.Mask.String(),
		FirstIPAvailable: NextIP(&ipv4Net.IP, 1).String(),
		LastIPAvailable:  NextIP(&ipv4Net.IP, IPAvailable).String(),
	}

	return info, nil
}

// NextIP - get next IP address
func NextIP(ip *net.IP, n int) *net.IP {
	newIP := ip.To4()

	for i := 0; i < n; i++ {
		newIP[3]++
	}

	return &newIP
}
