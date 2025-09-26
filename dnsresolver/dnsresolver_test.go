package main

import (
	"encoding/json"
	"testing"
)

func TestDNSResolver(t *testing.T) {
	got, err := DNSResolver("en.wikipedia.org", "", "8.8.8.8")
	expected := Result{
		Domain:    "en.wikipedia.org",
		Type:      "",
		DNSServer: "8.8.8.8",
		Answer:    []string{"2001:df2:e500:ed1a::1", "103.102.166.224"},
	}
	if err != nil {
		t.Errorf("expected result but got error: %v", err)
	}

	data, _ := json.MarshalIndent(expected, "", "   ")
	if got != string(data) {
		t.Errorf("expected: %s but got: %s", data, got)
	}
}
