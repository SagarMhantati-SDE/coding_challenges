package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func Dial(port *string, timeout *int) error {
	address := os.Args[len(os.Args)-1]
	fmt.Println("port:", *port)
	server := fmt.Sprintf("%s:%d", address, *port)
	fmt.Println("Server:", server)
	_, err := net.DialTimeout("tcp", server, time.Duration(*timeout)*time.Second)
	if err != nil {
		return fmt.Errorf("PORT_CLOSED")
	}
	return nil
}

func main() {

	port := flag.String("p", "80", `Ports or ranges (default "1-1024") Example: "22", "80,443", "22-25,80,443"`)
	flag.StringVar(port, "port", "80", `Ports or ranges (default "1-1024") Example: "22", "80,443", "22-25,80,443"`)

	timeout := flag.Int("t", 10, `Connection timeout per port in milliseconds (default 500)`)
	flag.IntVar(timeout, "timeout", 10, `Connection timeout per port in milliseconds (default 500)`)

	concurrency := flag.Int("c", 0, `Number of concurrent workers or goroutines (default 100)`)
	flag.IntVar(concurrency, "concurrency", 0, `Number of concurrent workers or goroutines (default 100)`)

	flag.Parse()

	// TODO: Add json result
	// TODO: Add concurrency code
	if strings.Contains(*port, ",") {
		ports := strings.Split(*port, ",")

		for _, value := range ports {
			if !strings.Contains(value, "-") {
				Dial(&value, timeout)
			} else {
				startPort := strings.Split(value, "-")[0]
				endPort := strings.Split(value, "-")[1]

				startPortInt, _ := strconv.ParseInt(startPort, 10, 10)
				endPortInt, _ := strconv.ParseInt(endPort, 10, 10)

				for i := startPortInt; i <= endPortInt; i++ {
					ii := fmt.Sprintf("%d", i)
					Dial(&ii, timeout)
				}
			}
		}
	}

	// err := Dial(port, timeout)
	// if err != nil {
	// 	fmt.Println("Port is open")
	// }
}
