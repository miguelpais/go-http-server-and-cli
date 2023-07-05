package routing

import "testing"

var cases = []struct {
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

func TestGetPathRequest(t *testing.T) {
	for _, tc := range cases {
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
