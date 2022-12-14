package parser

import (
	"regexp"
)

type EnvoyAccessLog struct {
	StartTime                      string `json:"start_time,omitempty" ltsv:"start_time"`
	Method                         string `json:"method,omitempty" ltsv:"method"`
	Path                           string `json:"path,omitempty" ltsv:"path"`
	Protocol                       string `json:"protocol,omitempty" ltsv:"protocol"`
	ResponseCode                   string `json:"response_code,omitempty" ltsv:"response_code"`
	ResponseFlags                  string `json:"response_flags,omitempty" ltsv:"response_flags"`
	ResponseCodeDetails            string `json:"response_code_details,omitempty" ltsv:"response_code_details"`
	ConnectionTerminationDetails   string `json:"connection_termination_details,omitempty" ltsv:"connection_termination_details"`
	UpstreamTransportFailureReason string `json:"upstream_transport_failure_reason,omitempty" ltsv:"upstream_transport_failure_reason"`
	BytesReceived                  string `json:"bytes_received,omitempty" ltsv:"bytes_received"`
	BytesSent                      string `json:"bytes_sent,omitempty" ltsv:"bytes_sent"`
	Duration                       string `json:"duration,omitempty" ltsv:"duration"`
	XEnvoyUpstreamServiceTime      string `json:"x_envoy_upstream_service_time,omitempty" ltsv:"x_envoy_upstream_service_time"`
	XForwardedFor                  string `json:"x_forwarded_for,omitempty" ltsv:"x_forwarded_for"`
	UserAgent                      string `json:"user_agent,omitempty" ltsv:"user_agent"`
	XRequestId                     string `json:"x_request_id,omitempty" ltsv:"x_request_id"`
	Authority                      string `json:"authority,omitempty" ltsv:"authority"`
	UpstreamHost                   string `json:"upstream_host,omitempty" ltsv:"upstream_host"`
	UpstreamCluster                string `json:"upstream_cluster,omitempty" ltsv:"upstream_cluster"`
	UpstreamLocalAddress           string `json:"upstream_local_address,omitempty" ltsv:"upstream_local_address"`
	DownstreamLocalAddress         string `json:"downstream_local_address,omitempty" ltsv:"downstream_local_address"`
	DownstreamRemoteAddress        string `json:"downstream_remote_address,omitempty" ltsv:"downstream_remote_address"`
	RequestedServerName            string `json:"requested_server_name,omitempty" ltsv:"requested_server_name"`
	RouteName                      string `json:"route_name,omitempty" ltsv:"route_name"`
}

// Logs are parsed based on Default asccess log format
// https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format
var logRe = regexp.MustCompile(
	`\[(\d{4}(?:.\d{2}){2}(?:\s|T)(?:\d{2}.){3}\d+.?)\]\s` + // [%START_TIME%]
		`\"(\S+)\s(\S+)\s(\S+)\"\s` + // \"%REQ(:METHOD)% %REQ(X-ENVOY-ORIGINAL-PATH?:PATH)% %PROTOCOL%\"
		`(\d+)\s` + // %RESPONSE_CODE%
		`(\S+)\s` + // %RESPONSE_FLAGS%
		`(\S+)\s` + // %RESPONSE_CODE_DETAILS%
		`(\S+)\s` + // %CONNECTION_TERMINATION_DETAILS%
		`\"(\S+)\"\s` + // \"%UPSTREAM_TRANSPORT_FAILURE_REASON%\"
		`(\d+)\s` + // %BYTES_RECEIVED%
		`(\d+)\s` + // %BYTES_SENT%
		`(\d+)\s` + // %DURATION%
		`(\d+)\s` + // %RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%
		`\"(\S+)\"\s` + // \"%REQ(X-FORWARDED-FOR)%\"
		`\"(.+)\"\s` + // \"%REQ(USER-AGENT)%\"
		`\"(\S+)\"\s` + // \"%REQ(X-REQUEST-ID)%\"
		`\"(\S+)\"\s` + // \"%REQ(:AUTHORITY)%\"
		`\"((?:\d{1,3}.){3}\d{1,3}:\d+)\"\s` + // \"%UPSTREAM_HOST%\"
		`(\S+)\s` + // %UPSTREAM_CLUSTER%
		`((?:\d{1,3}.){3}\d{1,3}:\d+)\s` + // %UPSTREAM_LOCAL_ADDRESS%
		`((?:\d{1,3}.){3}\d{1,3}:\d+)\s` + // %DOWNSTREAM_LOCAL_ADDRESS%
		`((?:\d{1,3}.){3}\d{1,3}:\d+)\s` + // %DOWNSTREAM_REMOTE_ADDRESS%
		`(\S+)\s` + // %REQUESTED_SERVER_NAME%
		`(\S+)$`) // %ROUTE_NAME%

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(line string) (*EnvoyAccessLog, error) {
	matches := logRe.FindStringSubmatch(line)
	if len(matches) < 1 {
		return &EnvoyAccessLog{}, nil
	}
	return &EnvoyAccessLog{
		StartTime:                      matches[1],
		Method:                         matches[2],
		Path:                           matches[3],
		Protocol:                       matches[4],
		ResponseCode:                   matches[5],
		ResponseFlags:                  matches[6],
		ResponseCodeDetails:            matches[7],
		ConnectionTerminationDetails:   matches[8],
		UpstreamTransportFailureReason: matches[9],
		BytesReceived:                  matches[10],
		BytesSent:                      matches[11],
		Duration:                       matches[12],
		XEnvoyUpstreamServiceTime:      matches[13],
		XForwardedFor:                  matches[14],
		UserAgent:                      matches[15],
		XRequestId:                     matches[16],
		Authority:                      matches[17],
		UpstreamHost:                   matches[18],
		UpstreamCluster:                matches[19],
		UpstreamLocalAddress:           matches[20],
		DownstreamLocalAddress:         matches[21],
		DownstreamRemoteAddress:        matches[22],
		RequestedServerName:            matches[23],
		RouteName:                      matches[24],
	}, nil
}
