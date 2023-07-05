package routing

import (
	"testing"
)

var positiveScenarios = []struct {
	desc     string
	in       string
	expected string
}{
	{"ui root index",
		"GET / HTTP/1.1\r\n" +
			"Host: localhost:8000\r\n" +
			"User-Agent: Mozilla/5.0\r\n" +
			"Accept: */*\r\n",
		"/"},
	{"api index",
		"GET /api HTTP/1.1\r\n" +
			"Host: localhost:8000\r\n" +
			"User-Agent: Mozilla/5.0\r\n" +
			"Accept: */*\r\n",
		"/api"},
}

func TestGetPathForRequestPositiveScenarios(t *testing.T) {
	for _, tc := range positiveScenarios {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := GetRequestPath(tc.in)
			if err != nil {
				t.Fatalf("Could not parse the incoming request, error was %s", err)
			}
			if result != tc.expected {
				t.Fatalf("Parsing result for incoming request was not %s, instead was %s", tc.expected, result)
			}
		})
	}
}

var invalidScenarios = []struct {
	desc string
	in   string
}{
	{"blank string received",
		""},
	{"no path in http request ",
		"GET HTTP/1.1\r\n"},
	{"double path http request ",
		"GET / / HTTP/1.1\r\n"},
	{"path does not start with / http request ",
		"GET abc HTTP/1.1\r\n"},
	{"giberish request",
		"this is a lot of giberish"},
}

func TestGetPathForRequestInvalidScenarios(t *testing.T) {
	for _, tc := range invalidScenarios {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := GetRequestPath(tc.in)
			if err == nil {
				t.Fatalf("Test case did not return expected err, it returne instead %s", result)
			}
		})
	}
}
