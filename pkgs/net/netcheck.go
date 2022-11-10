package netcheck

import (
	"context"
	"fmt"

	// "io"
	"net"
	// "net/http"
	// "net/http/httptest"
	"os"
	"time"

	"github.com/go-ping/ping"
	netcheker "github.com/tevino/tcp-shaker"
	"golang.org/x/net/nettest"
)

// check Support system of IPv4:
func IsSupportIP4() (ok bool) {
	isOk := nettest.SupportsIPv4()
	return isOk
}

// check what ports is open:
func OpenPorts(domain string) []string {
	ports := [...]string{"80", "88", "443", "8080", "8888", "28800", "28888", "50808", "65480"}
	checkedPorts := []string{}
	for i := 0; i < len(ports); i++ {
		port := ports[i]
		conn, err := net.Dial("tcp", domain+":"+port)
		if err == nil {
			checkedPorts = append([]string(checkedPorts), port)
			conn.Close()
		}
	}
	return checkedPorts
}

// ping to check any domain or ip:
func PingCheck(domain string) (ok bool) {
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	// pinger.Timeout = 500

	// fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	ok = true
	return ok
}

// check reachable website:
func NetChecker(host string) {
	c := netcheker.NewChecker()

	ctx, stopChecker := context.WithCancel(context.Background())
	defer stopChecker()
	go func() {
		if err := c.CheckingLoop(ctx); err != nil {
			fmt.Println("checking loop stopped due to fatal error: ", err)
		}
	}()

	<-c.WaitReady()

	timeout := time.Second * 1
	err := c.CheckAddr(host+":80", timeout)
	switch err {
	case netcheker.ErrTimeout:
		fmt.Println("Connect to " + host + " timed out")
	case nil:
		fmt.Println("Connect to " + host + " succeeded")
	default:
		fmt.Println("Error occurred while connecting: ", err)
	}
}

func IsDNSError(err error) bool {
	if _, ok := err.(*net.DNSError); ok {
		return true
	}
	return false
}

// check what Interfaces is available:
func AvailableInterfaces() {
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, i := range interfaces {
		fmt.Printf("Name : %v \n", i.Name)
	}
}

// func ResponseRecorder(domain string) {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Connected")
// 	}

// 	req := httptest.NewRequest("GET", domain, nil)
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	resp := w.Result()
// 	body, _ := io.ReadAll(resp.Body)

// 	if resp.StatusCode == 200 {
// 		fmt.Println("Connected to " + domain + " is OK")
// 	}

// 	// fmt.Println(resp.StatusCode)
// 	// fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))

// }

// func isTCPWorking(c net.Conn) bool {
// 	_, err := c.Write([]byte("OK"))
// 	if err != nil {
// 		c.Close() // close if problem
// 		return false
// 	}
// 	return true
// }

// func portForServices(service string) {

// 	port, err := net.LookupPort("tcp", service)
// 	if err == nil {
// 		fmt.Println(service+" port: ", port)
// 	}
// }

// func systemInterfaces() {
// 	interfaces, _ := net.Interfaces()
// 	for _, inter := range interfaces {
// 		fmt.Println("Index :", inter.Index)
// 		fmt.Println("Name  :", inter.Name)
// 		fmt.Println("HWaddr:", inter.HardwareAddr)
// 		fmt.Println("MTU   :", inter.MTU)
// 		fmt.Println("Flags :", inter.Flags)
// 		addrs, _ := inter.Addrs()
// 		for _, ipaddr := range addrs {
// 			fmt.Println("Addr  :", ipaddr)
// 		}
// 		fmt.Println()
// 	}
// }
