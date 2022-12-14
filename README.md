# istio-axslog

[![build](https://github.com/rnakamine/istio-axslog/workflows/build/badge.svg?branch=main)](https://github.com/rnakamine/istio-axslog/actions?workflow=build)
[![MITLicense](https://img.shields.io/github/license/rnakamine/istio-axslog)](https://github.com/rnakamine/istio-axslog/blob/main/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rnakamine/istio-axslog)](https://pkg.go.dev/github.com/rnakamine/istio-axslog)

istio-axslog is parsed istio-proxy(envoy) access log and output in any format. Logs are parsed based on [Istio / Default access log format](https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format).

## Install

### homebrew

```sh
$ brew install rnakamine/tap/istio-axslog
```

### manually

Download binary from [releases page](https://github.com/rnakamine/istio-axslog/releases)


## Usage
```sh
$ istio-axslog --help
istio-axslog is parsed istio-proxy(envoy) access log and output in any format.

Usage:
  istio-axslog [flags]

Flags:
  -h, --help            help for istio-axslog
  -o, --output string   output format. supported formats are json, ltsv (default "json")
  -v, --version         version for istio-axslog
```
Logs can be received from STDIN. Also, the supported output formats are `json`, `ltsv`.

#### Output in json format (default):

```sh
# istio-axslog (--output json)
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default' | istio-axslog
{"start_time":"2020-11-25T21:26:18.409Z","method":"GET","path":"/status/418","protocol":"HTTP/1.1","response_code":"418","response_flags":"-","response_code_details":"via_upstream","connection_termination_details":"-","upstream_transport_failure_reason":"-","bytes_received":"0","bytes_sent":"135","duration":"4","x_envoy_upstream_service_time":"4","x_forwarded_for":"-","user_agent":"curl/7.73.0-DEV","x_request_id":"84961386-6d84-929d-98bd-c5aee93b5c88","authority":"httpbin:8000","upstream_host":"10.44.1.27:80","upstream_cluster":"outbound|8000||httpbin.foo.svc.cluster.local","upstream_local_address":"10.44.1.23:37652","downstream_local_address":"10.0.45.184:8000","downstream_remote_address":"10.44.1.23:46520","requested_server_name":"-","route_name":"default"}
```

#### Output in ltsv format:
```sh
# istio-axslog --output ltsv
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default' | istio-axslog --output ltsv
start_time:2020-11-25T21:26:18.409Z     method:GET      path:/status/418        protocol:HTTP/1.1       response_code:418       response_flags:-        response_code_details:via_upstream      connection_termination_details:-        upstream_transport_failure_reason:-     bytes_received:0        bytes_sent:135     duration:4      x_envoy_upstream_service_time:4 x_forwarded_for:-       user_agent:curl/7.73.0-DEV      x_request_id:84961386-6d84-929d-98bd-c5aee93b5c88       authority:httpbin:8000  upstream_host:10.44.1.27:80     upstream_cluster:outbound|8000||httpbin.foo.svc.cluster.local   upstream_local_address:10.44.1.23:37652    downstream_local_address:10.0.45.184:8000       downstream_remote_address:10.44.1.23:46520      requested_server_name:- route_name:default
```
