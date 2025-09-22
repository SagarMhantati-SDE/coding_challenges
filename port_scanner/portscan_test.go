package main

import "testing"

func Compare(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {
		if value != a[key] {
			return false
		}
	}

	return true
}
func TestParsePorts(t *testing.T) {

	testcases := []string{
		"10,11,12",
		"10-12,15",
		"11,12,13,19,20-25,27",
	}

	expected := [][]string{
		{"10", "11", "12"},
		{"10", "11", "12", "15"},
		{"11", "12", "13", "19", "20", "21", "22", "23", "24", "25", "27"},
	}

	for key, value := range testcases {
		got := ParsePorts(value)
		if !Compare(got, expected[key]) {
			t.Errorf("expected: %v, got: %v", expected[key], got)
		}
	}
}
