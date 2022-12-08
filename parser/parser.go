package parser

import (
	"github.com/vjeantet/grok"
)

// ref: https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format
var defaultAccessLogFormat = `\[%{TIMESTAMP_ISO8601:start_time}\] ` +
	`\"%{DATA:method} %{DATA:path} %{DATA:protocol}\" ` +
	`%{NUMBER:response_code} %{DATA:response_flags} ` +
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

func (p *Parser) Parse(line string) (map[string]string, error) {
	m, err := p.grok.Parse(defaultAccessLogFormat, line)
	if err != nil {
		return nil, err
	}

	return m, nil
}
