package cidr

// IPAddressInfo - IP adress infomation
type IPAddressInfo struct {
	IPAvailable      int
	Mask             string
	FirstIPAvailable string
	LastIPAvailable  string
	BoardCase        string
}

// ParseCIDR - Parse the given CIDR to IP adress infomation
func ParseCIDR() (IPAddressInfo, error) {
	return IPAddressInfo{}, nil
}
