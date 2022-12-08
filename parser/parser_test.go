package parser

import (
	"reflect"
	"testing"
)

var parseTests = []struct {
	name  string
	input string
	want  map[string]string
}{
	{
		name:  "log from sleep",
		input: `[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default`,
		want: map[string]string{
			"bytes_received":                    "0",
			"duration":                          "4",
			"x-forwarded-for":                   "-",
			"user-agent":                        "curl/7.73.0-DEV",
			"upstream_local_address":            "10.44.1.23:37652",
			"downstream_local_address":          "10.0.45.184:8000",
			"response_code":                     "418",
			"connection_termination_details":    "-",
			"downstream_remote_address":         "10.44.1.23:46520",
			"requested_server_name":             "-",
			"path":                              "/status/418",
			"response_code_details":             "via_upstream",
			"bytes_sent":                        "135",
			"upstream_host":                     "10.44.1.27:80",
			"start_time":                        "2020-11-25T21:26:18.409Z",
			"method":                            "GET",
			"route_name":                        "default",
			"response_flags":                    "-",
			"authority":                         "httpbin:8000",
			"x-envoy-upstream-service-time":     "4",
			"x-request-id":                      "84961386-6d84-929d-98bd-c5aee93b5c88",
			"upstream_cluster":                  "outbound|8000||httpbin.foo.svc.cluster.local",
			"protocol":                          "HTTP/1.1",
			"upstream_transport_failure_reason": "-",
		},
	},
	{
		name:  "log from httpbin",
		input: `[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 3 1 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "127.0.0.1:80" inbound|8000|| 127.0.0.1:41854 10.44.1.27:80 10.44.1.23:37652 outbound_.8000_._.httpbin.foo.svc.cluster.local default`,
		want: map[string]string{
			"bytes_received":                    "0",
			"bytes_sent":                        "135",
			"duration":                          "3",
			"x-forwarded-for":                   "-",
			"requested_server_name":             "outbound_.8000_._.httpbin.foo.svc.cluster.local",
			"response_flags":                    "-",
			"upstream_cluster":                  "inbound|8000||",
			"downstream_remote_address":         "10.44.1.23:37652",
			"route_name":                        "default",
			"start_time":                        "2020-11-25T21:26:18.409Z",
			"method":                            "GET",
			"response_code_details":             "via_upstream",
			"x-envoy-upstream-service-time":     "1",
			"user-agent":                        "curl/7.73.0-DEV",
			"x-request-id":                      "84961386-6d84-929d-98bd-c5aee93b5c88",
			"path":                              "/status/418",
			"protocol":                          "HTTP/1.1",
			"response_code":                     "418",
			"connection_termination_details":    "-",
			"upstream_transport_failure_reason": "-",
			"authority":                         "httpbin:8000",
			"upstream_host":                     "127.0.0.1:80",
			"upstream_local_address":            "127.0.0.1:41854",
			"downstream_local_address":          "10.44.1.27:80",
		},
	},
}

func TestParse(t *testing.T) {
	p, _ := New()
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := p.Parse(tt.input)
			if !reflect.DeepEqual(m, tt.want) {
				t.Errorf("Parse() = %v, want %v", m, tt.want)
			}
		})
	}
}
