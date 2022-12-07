package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/vjeantet/grok"
)

type EnvoyAccessLog struct {
	Authority                    string `mapstructure:"authority" json:"authority,omitempty"`
	BytesReceived                string `mapstructure:"bytes_received" json:"bytes_received,omitempty"`
	BytesSent                    string `mapstructure:"bytes_sent" json:"bytes_sent,omitempty"`
	ConnectionTerminationDetails string `mapstructure:"termination_details" json:"connection_termination_details,omitempty"`
	Duration                     string `mapstructure:"duration" json:"duration,omitempty"`
	ForwardedFor                 string `mapstructure:"forwarded_for" json:"forwarded_for,omitempty"`
	Method                       string `mapstructure:"method" json:"method,omitempty"`
	Protocol                     string `mapstructure:"protocol" json:"protocol,omitempty"`
	RequestId                    string `mapstructure:"request_id" json:"request_id,omitempty"`
	ResponseFlags                string `mapstructure:"response_flags" json:"response_flags,omitempty"`
	StatusCode                   string `mapstructure:"status_code" json:"status_code,omitempty"`
	TcpServiceTime               string `mapstructure:"tcp_service_time" json:"tcp_service_time,omitempty"`
	Timestamp                    string `mapstructure:"timestamp" json:"timestamp,omitempty"`
	UpstreamService              string `mapstructure:"upstream_service" json:"upstream_service,omitempty"`
	UpstreamServiceTime          string `mapstructure:"upstream_service_time" json:"upstream_service_time,omitempty"`
	UpstreamCluster              string `mapstructure:"upstream_cluster" json:"upstream_cluster,omitempty"`
	UpstreamLocal                string `mapstructure:"upstream_local" json:"upstream_local,omitempty"`
	DownstreamLocal              string `mapstructure:"downstream_local" json:"downstream_local,omitempty"`
	DownstreamRemote             string `mapstructure:"downstream_remote" json:"downstream_remote,omitempty"`
	RequestedServer              string `mapstructure:"requested_server" json:"requested_server,omitempty"`
	ResponseCodeDetails          string `mapstructure:"response_details" json:"response_code_details,omitempty"`
	RouteName                    string `mapstructure:"route_name" json:"route_name,omitempty"`
	UpstreamFailureReason        string `mapstructure:"upstream_failure_reason" json:"upstream_failure_reason,omitempty"`
	UriParam                     string `mapstructure:"uri_param" json:"uri_param,omitempty"`
	UriPath                      string `mapstructure:"uri_path" json:"uri_path,omitempty"`
	UserAgent                    string `mapstructure:"user_agent" json:"user_agent,omitempty"`
}

const envoyAccessLogPattern string = `\[%{TIMESTAMP_ISO8601:timestamp}\] \"%{DATA:method} (?:(?:%{URIPATH:uri_path}(?:%{URIPARAM:uri_param})?)|%{DATA}) %{DATA:protocol}\" %{NUMBER:status_code} %{DATA:response_flags} %{DATA:response_details} %{DATA:termination_details} \"%{DATA:upstream_failure_reason}\" %{NUMBER:bytes_received} %{NUMBER:bytes_sent} %{NUMBER:duration} (?:%{NUMBER:upstream_service_time}|%{DATA:tcp_service_time}) \"%{DATA:forwarded_for}\" \"%{DATA:user_agent}\" \"%{DATA:request_id}\" \"%{DATA:authority}\" \"%{DATA:upstream_service}\" %{DATA:upstream_cluster} %{DATA:upstream_local} %{DATA:downstream_local} %{DATA:downstream_remote} %{DATA:requested_server}(?: %{DATA:route_name})?$`

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

func (p *Parser) Parse(log string) (*EnvoyAccessLog, error) {
	m, err := p.grok.Parse(envoyAccessLogPattern, log)
	if err != nil {
		return nil, err
	}

	accessLog := &EnvoyAccessLog{}
	err = mapstructure.Decode(m, accessLog)
	if err != nil {
		return nil, err
	}

	return accessLog, nil
}
