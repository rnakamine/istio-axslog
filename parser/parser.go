package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/vjeantet/grok"
)

type EnvoyAccessLog struct {
	StartTime                      string `mapstructure:"start_time" json:"start_time,omitempty"`
	Method                         string `mapstructure:"method" json:"method,omitempty"`
	Path                           string `mapstructure:"path" json:"path,omitempty"`
	Protocol                       string `mapstructure:"protocol" json:"protocol,omitempty"`
	ResponseCode                   string `mapstructure:"response_code" json:"response_code,omitempty"`
	ResponseFlags                  string `mapstructure:"response_flags" json:"response_flags,omitempty"`
	ResponseCodeDetails            string `mapstructure:"response_code_details" json:"response_code_details,omitempty"`
	ConnectionTerminationDetails   string `mapstructure:"connection_termination_details" json:"connection_termination_details,omitempty"`
	UpstreamTransportFailureReason string `mapstructure:"upstream_transport_failure_reason" json:"upstream_transport_failure_reason,omitempty"`
	BytesReceived                  string `mapstructure:"bytes_received" json:"bytes_received,omitempty"`
	BytesSent                      string `mapstructure:"bytes_sent" json:"bytes_sent,omitempty"`
	Duration                       string `mapstructure:"duration" json:"duration,omitempty"`
	XEnvoyUpstreamServiceTime      string `mapstructure:"x-envoy-upstream-service-time" json:"x-envoy-upstream-service-time,omitempty"`
	XForwardedFor                  string `mapstructure:"x-forwarded-for" json:"x-forwarded-for,omitempty"`
	UserAgent                      string `mapstructure:"user-agent" json:"user-agent,omitempty"`
	XRequestId                     string `mapstructure:"x-request-id" json:"x-request-id,omitempty"`
	Authority                      string `mapstructure:"authority" json:"authority,omitempty"`
	UpstreamHost                   string `mapstructure:"upstream_host" json:"upstream_host,omitempty"`
	UpstreamCluster                string `mapstructure:"upstream_cluster" json:"upstream_cluster,omitempty"`
	UpstreamLocalAddress           string `mapstructure:"upstream_local_address" json:"upstream_local_address,omitempty"`
	DownstreamLocalAddress         string `mapstructure:"downstream_local_address" json:"downstream_local_address,omitempty"`
	DownstreamRemoteAddress        string `mapstructure:"downstream_remote_address" json:"downstream_remote_address,omitempty"`
	RequestedServerName            string `mapstructure:"requested_server_name" json:"requested_server_name,omitempty"`
	RouteName                      string `mapstructure:"route_name" json:"route_name,omitempty"`
}

// ref: https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format
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
	if len(m) < 1 {
		return nil, nil
	}
	accessLog := &EnvoyAccessLog{}
	err = mapstructure.Decode(m, accessLog)
	if err != nil {
		return accessLog, err
	}
	return accessLog, nil
}
