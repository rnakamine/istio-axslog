package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/vjeantet/grok"
)

type EnvoyAccessLog struct {
	StartTime                      string `mapstructure:"start_time" json:"start_time,omitempty" ltsv:"start_time"`
	Method                         string `mapstructure:"method" json:"method,omitempty" ltsv:"method"`
	Path                           string `mapstructure:"path" json:"path,omitempty" ltsv:"path"`
	Protocol                       string `mapstructure:"protocol" json:"protocol,omitempty" ltsv:"protocol"`
	ResponseCode                   string `mapstructure:"response_code" json:"response_code,omitempty" ltsv:"response_code"`
	ResponseFlags                  string `mapstructure:"response_flags" json:"response_flags,omitempty" ltsv:"response_flags"`
	ResponseCodeDetails            string `mapstructure:"response_code_details" json:"response_code_details,omitempty" ltsv:"response_code_details"`
	ConnectionTerminationDetails   string `mapstructure:"connection_termination_details" json:"connection_termination_details,omitempty" ltsv:"connection_termination_details"`
	UpstreamTransportFailureReason string `mapstructure:"upstream_transport_failure_reason" json:"upstream_transport_failure_reason,omitempty" ltsv:"upstream_transport_failure_reason"`
	BytesReceived                  string `mapstructure:"bytes_received" json:"bytes_received,omitempty" ltsv:"bytes_received"`
	BytesSent                      string `mapstructure:"bytes_sent" json:"bytes_sent,omitempty" ltsv:"bytes_sent"`
	Duration                       string `mapstructure:"duration" json:"duration,omitempty" ltsv:"duration"`
	XEnvoyUpstreamServiceTime      string `mapstructure:"x-envoy-upstream-service-time" json:"x-envoy-upstream-service-time,omitempty" ltsv:"x-envoy-upstream-service-time"`
	XForwardedFor                  string `mapstructure:"x-forwarded-for" json:"x-forwarded-for,omitempty" ltsv:"x-forwarded-for"`
	UserAgent                      string `mapstructure:"user-agent" json:"user-agent,omitempty" ltsv:"user-agent"`
	XRequestId                     string `mapstructure:"x-request-id" json:"x-request-id,omitempty" ltsv:"x-request-id"`
	Authority                      string `mapstructure:"authority" json:"authority,omitempty" ltsv:"authority"`
	UpstreamHost                   string `mapstructure:"upstream_host" json:"upstream_host,omitempty" ltsv:"upstream_host"`
	UpstreamCluster                string `mapstructure:"upstream_cluster" json:"upstream_cluster,omitempty" ltsv:"upstream_cluster"`
	UpstreamLocalAddress           string `mapstructure:"upstream_local_address" json:"upstream_local_address,omitempty" ltsv:"upstream_local_address"`
	DownstreamLocalAddress         string `mapstructure:"downstream_local_address" json:"downstream_local_address,omitempty" ltsv:"downstream_local_address"`
	DownstreamRemoteAddress        string `mapstructure:"downstream_remote_address" json:"downstream_remote_address,omitempty" ltsv:"downstream_remote_address"`
	RequestedServerName            string `mapstructure:"requested_server_name" json:"requested_server_name,omitempty" ltsv:"requested_server_name"`
	RouteName                      string `mapstructure:"route_name" json:"route_name,omitempty" ltsv:"route_name"`
}

// Logs are parsed based on Default asccess log format
// https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format
var defaultAccessLogFormat = `\[%{TIMESTAMP_ISO8601:start_time}\] ` +
	`\"%{DATA:method} %{DATA:path} %{DATA:protocol}\" ` +
	`%{NUMBER:response_code} ` +
	`%{DATA:response_flags} ` +
	`%{DATA:response_code_details} ` +
	`%{DATA:connection_termination_details} ` +
	`\"%{DATA:upstream_transport_failure_reason}\" ` +
	`%{NUMBER:bytes_received} ` +
	`%{NUMBER:bytes_sent} ` +
	`%{NUMBER:duration} ` +
	`%{NUMBER:x-envoy-upstream-service-time} ` +
	`\"%{DATA:x-forwarded-for}\" ` +
	`\"%{DATA:user-agent}\" ` +
	`\"%{DATA:x-request-id}\" ` +
	`\"%{DATA:authority}\" ` +
	`\"%{DATA:upstream_host}\" ` +
	`%{DATA:upstream_cluster} ` +
	`%{DATA:upstream_local_address} ` +
	`%{DATA:downstream_local_address} ` +
	`%{DATA:downstream_remote_address} ` +
	`%{DATA:requested_server_name} ` +
	`%{DATA:route_name}$`

type Parser struct {
	grok *grok.Grok
}

func New() (*Parser, error) {
	g, err := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	if err != nil {
		return nil, err
	}
	return &Parser{g}, nil
}

func (p *Parser) Parse(line string) (*EnvoyAccessLog, error) {
	m, err := p.grok.Parse(defaultAccessLogFormat, line)
	if err != nil {
		return nil, err
	}
	accessLog := &EnvoyAccessLog{}
	if len(m) < 1 {
		return accessLog, nil
	}
	err = mapstructure.Decode(m, accessLog)
	if err != nil {
		return nil, err
	}
	return accessLog, nil
}
