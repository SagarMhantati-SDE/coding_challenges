package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
)

type Result struct {
	Domain    string   `json:"domain"`
	Type      string   `json:"type"`
	DNSServer string   `json:"dns_server"`
	Answer    []string `json:"answer"`
}

func DNSResolver(domain, dtype, server string) (string, error) {
	resolver := net.Resolver{}
	ctx := context.WithValue(context.Background(), "host", server)
	addrs, err := resolver.LookupHost(ctx, domain)
	if err != nil {
		return "", fmt.Errorf("error occurred during dns resolution, error: %v", err)
	}

	result := Result{
		Domain:    domain,
		Type:      dtype,
		DNSServer: server,
		Answer:    addrs,
	}
	data, err := json.MarshalIndent(result, "", "   ")

	if err != nil {
		return "", fmt.Errorf("error occurred during json encoding, error: %v", err)
	}

	fmt.Printf("%v", string(data))
	return string(data), nil
}

func main() {
	var domain string
	var dtype string
	var server string

	flag.StringVar(&domain, "domain", "example.com", "domain name")
	flag.StringVar(&dtype, "dtype", "A", "domain type")
	flag.StringVar(&server, "server", "8.8.8.8", "server of domain name")

	flag.Parse()

	if _, err := DNSResolver(domain, dtype, server); err != nil {
		fmt.Println("Unable resolve the dns")
	}
}
