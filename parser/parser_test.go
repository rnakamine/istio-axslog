package parser

import (
	"testing"
)

var parseTests = []struct {
	name  string
	input string
	want  *EnvoyAccessLog
}{
	{
		name:  "log from sleep",
		input: `[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default`,
		want: &EnvoyAccessLog{
			StartTime:                      "2020-11-25T21:26:18.409Z",
			Method:                         "GET",
			Path:                           "/status/418",
			Protocol:                       "HTTP/1.1",
			ResponseCode:                   "418",
			ResponseFlags:                  "-",
			ResponseCodeDetails:            "via_upstream",
			ConnectionTerminationDetails:   "-",
			UpstreamTransportFailureReason: "-",
			BytesReceived:                  "0",
			BytesSent:                      "135",
			Duration:                       "4",
			XEnvoyUpstreamServiceTime:      "4",
			XForwardedFor:                  "-",
			UserAgent:                      "curl/7.73.0-DEV",
			XRequestId:                     "84961386-6d84-929d-98bd-c5aee93b5c88",
			Authority:                      "httpbin:8000",
			UpstreamHost:                   "10.44.1.27:80",
			UpstreamCluster:                "outbound|8000||httpbin.foo.svc.cluster.local",
			UpstreamLocalAddress:           "10.44.1.23:37652",
			DownstreamLocalAddress:         "10.0.45.184:8000",
			DownstreamRemoteAddress:        "10.44.1.23:46520",
			RequestedServerName:            "-",
			RouteName:                      "default",
		},
	},
	{
		name:  "log from httpbin",
		input: `[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 3 1 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "127.0.0.1:80" inbound|8000|| 127.0.0.1:41854 10.44.1.27:80 10.44.1.23:37652 outbound_.8000_._.httpbin.foo.svc.cluster.local default`,
		want: &EnvoyAccessLog{
			StartTime:                      "2020-11-25T21:26:18.409Z",
			Method:                         "GET",
			Path:                           "/status/418",
			Protocol:                       "HTTP/1.1",
			ResponseCode:                   "418",
			ResponseFlags:                  "-",
			ResponseCodeDetails:            "via_upstream",
			ConnectionTerminationDetails:   "-",
			UpstreamTransportFailureReason: "-",
			BytesReceived:                  "0",
			BytesSent:                      "135",
			Duration:                       "3",
			XEnvoyUpstreamServiceTime:      "1",
			XForwardedFor:                  "-",
			UserAgent:                      "curl/7.73.0-DEV",
			XRequestId:                     "84961386-6d84-929d-98bd-c5aee93b5c88",
			Authority:                      "httpbin:8000",
			UpstreamHost:                   "127.0.0.1:80",
			UpstreamCluster:                "inbound|8000||",
			UpstreamLocalAddress:           "127.0.0.1:41854",
			DownstreamLocalAddress:         "10.44.1.27:80",
			DownstreamRemoteAddress:        "10.44.1.23:37652",
			RequestedServerName:            "outbound_.8000_._.httpbin.foo.svc.cluster.local",
			RouteName:                      "default",
		},
	},
	{
		name:  "log that do not match envoy's access log format",
		input: `[xxxxxxxxxx] "xxxxxxxxxx" xxxxxxxxxx`,
		want:  &EnvoyAccessLog{},
	},
}

func TestParse(t *testing.T) {
	p, _ := New()
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			accessLog, _ := p.Parse(tt.input)
			if *accessLog != *tt.want {
				t.Errorf("Parse() = %v, want %v", *accessLog, tt.want)
			}
		})
	}
}
