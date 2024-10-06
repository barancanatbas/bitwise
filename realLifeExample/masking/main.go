package main

import (
	"fmt"
	"net"
)

func subnetMasking(ipStr, maskStr string) (net.IP, error) {
	ip := net.ParseIP(ipStr).To4()
	mask := net.IPMask(net.ParseIP(maskStr).To4())

	network := net.IP(make([]byte, 4))
	for i := 0; i < 4; i++ {
		network[i] = ip[i] & mask[i]
	}
	return network, nil
}

func main() {
	ip := "192.168.1.10"
	mask := "255.255.255.0"

	network, err := subnetMasking(ip, mask)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Network address for %s with mask %s is %s\n", ip, mask, network)
}
