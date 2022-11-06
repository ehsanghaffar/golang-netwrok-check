package main

import (
	"fmt"
	"strings"

	netcheck "goChecker/pkgs/net"
)

func main() {

	// check ipv4
	ipv4 := netcheck.IsSupportIP4()
	if ipv4 == false {
		fmt.Println("This system not support IPv4")
		panic(ipv4)
	} else {
		fmt.Println("IPv4 support in this system")
	}

	// Check ports
	checkedPort := netcheck.OpenPorts("127.0.0.1")
	fmt.Println("What Oxygate's Port list open on this pc: ", strings.Join(checkedPort, " - "))

	// ping iranian DNS
	checkIranDNS := netcheck.PingCheck("178.22.122.100")
	if checkIranDNS == false {
		fmt.Println("You haven't هیجی")
		panic(checkIranDNS)
	} else {
		fmt.Println("Pinged to Iran DNS?: ", checkIranDNS)
	}

	// Check Iranian Internet:
	netcheck.NetChecker("khomeini.ir")

	// ping Google DNS
	checkGoogleDNS := netcheck.PingCheck("8.8.4.4")
	if checkGoogleDNS == false {
		fmt.Println("You haven't هیجی")
		panic(checkGoogleDNS)
	} else {
		fmt.Println("Pinged to Google DNS?: ", checkGoogleDNS)
	}

	// Check international Intetnet
	netcheck.NetChecker("google.com")

	// check oxygate
	netcheck.NetChecker("oxygate.com")

}
