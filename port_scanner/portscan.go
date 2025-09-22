package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	PORT_CLOSED = -1
	PORT_OPEN   = 0
)

type Result struct {
	Target       string    `json:"string"`
	PortsScanned []string  `json:"ports_scanned"`
	OpenPorts    []string  `json:"open_ports"`
	ClosedPorts  []string  `json:"closed_ports"`
	ElaspedTime  string    `json:"elasped_time"`
	TimeStamp    time.Time `json:"timestamp"`
}

func Dial(address string, port *string, timeout *int, wg *sync.WaitGroup) int {
	defer wg.Done()
	fmt.Println("port:", *port)
	server := fmt.Sprintf("%s:%s", address, *port)
	fmt.Println("Server:", server)
	_, err := net.DialTimeout("tcp", server, time.Duration(*timeout)*time.Second)
	if err != nil {
		return PORT_CLOSED
	}
	return PORT_OPEN
}

func ParsePorts(inputdata string) []string {
	var allPorts []string

	if strings.Contains(inputdata, ",") {
		ports := strings.Split(inputdata, ",")

		for _, value := range ports {
			if !strings.Contains(value, "-") {
				allPorts = append(allPorts, value)
			} else {
				startPort := strings.Split(value, "-")[0]
				endPort := strings.Split(value, "-")[1]

				startPortInt, _ := strconv.ParseInt(startPort, 10, 10)
				endPortInt, _ := strconv.ParseInt(endPort, 10, 10)

				for i := startPortInt; i <= endPortInt; i++ {
					ii := fmt.Sprintf("%d", i)
					allPorts = append(allPorts, ii)
				}
			}
		}
	}
	return allPorts
}

func main() {

	port := flag.String("p", "80", `Ports or ranges (default "1-1024") Example: "22", "80,443", "22-25,80,443"`)
	flag.StringVar(port, "port", "80", `Ports or ranges (default "1-1024") Example: "22", "80,443", "22-25,80,443"`)

	timeout := flag.Int("t", 10, `Connection timeout per port in milliseconds (default 500)`)
	flag.IntVar(timeout, "timeout", 10, `Connection timeout per port in milliseconds (default 500)`)

	concurrency := flag.Int("c", 0, `Number of concurrent workers or goroutines (default 100)`)
	flag.IntVar(concurrency, "concurrency", 0, `Number of concurrent workers or goroutines (default 100)`)

	flag.Parse()

	// TODO: Add concurrency code
	var openPorts []string
	var closedPorts []string
	address := os.Args[len(os.Args)-1]
	allPorts := ParsePorts(*port)
	var wg sync.WaitGroup
	startTime := time.Now()
	for _, value := range allPorts {
		wg.Add(1)
		go func() {
			status := Dial(address, &value, timeout, &wg)
			if status == PORT_OPEN {
				openPorts = append(openPorts, value)
			} else if status == PORT_CLOSED {
				closedPorts = append(closedPorts, value)
			}
		}()
	}

	wg.Wait()
	endTime := time.Since(startTime)
	result := Result{
		Target:       address,
		PortsScanned: allPorts,
		OpenPorts:    openPorts,
		ClosedPorts:  closedPorts,
		ElaspedTime:  fmt.Sprintf("%ds", int64(endTime.Seconds())),
		TimeStamp:    time.Now(),
	}

	fmt.Printf("Elasped time: %s", endTime)
	data, _ := json.Marshal(result)
	fmt.Printf("data: %v", string(data))
}

// With go routine Elasped time: 10.001802834s
// Without go routine Elasped time: 40.032372542s
