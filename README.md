# istio-axslog

istio-axslog is parsed istio-proxy(envoy) access log and output in any format.
istio-axslog parses the istio-proxy(envoy) access log and outputs it in any format. Logs are parsed based on [Istio / Default access log format](https://istio.io/latest/docs/tasks/observability/logs/access-log/#default-access-log-format).

## Usage
```sh
$ istio-axslog --help
istio-axslog is parsed istio-proxy(envoy) access log and output in any format.

Usage:
  istio-axslog [flags]

Flags:
  -h, --help            help for istio-axslog
  -o, --output string   output format. supported formats are json, ltsv (default "json")
```
Logs can be received from STDIN. Also, the supported output formats are `json`, `ltsv`

#### Output in json format (default):

```sh
# istio-axslog (--output json)
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default' | istio-axslog
{"start_time":"2020-11-25T21:26:18.409Z","method":"GET","path":"/status/418","protocol":"HTTP/1.1","response_code":"418","response_flags":"-","response_code_details":"via_upstream","connection_termination_details":"-","upstream_transport_failure_reason":"-","bytes_received":"0","bytes_sent":"135","duration":"4","x-envoy-upstream-service-time":"4","x-forwarded-for":"-","user-agent":"curl/7.73.0-DEV","x-request-id":"84961386-6d84-929d-98bd-c5aee93b5c88","authority":"httpbin:8000","upstream_host":"10.44.1.27:80","upstream_cluster":"outbound|8000||httpbin.foo.svc.cluster.local","upstream_local_address":"10.44.1.23:37652","downstream_local_address":"10.0.45.184:8000","downstream_remote_address":"10.44.1.23:46520","requested_server_name":"-","route_name":"default"}

```

#### Output in ltsv format:
```sh
# istio-axslog --output ltsv
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 4 4 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "10.44.1.27:80" outbound|8000||httpbin.foo.svc.cluster.local 10.44.1.23:37652 10.0.45.184:8000 10.44.1.23:46520 - default' | istio-axslog --output ltsv
start_time:2020-11-25T21:26:18.409Z     method:GET      path:/status/418        protocol:HTTP/1.1       response_code:418       response_flags:-        response_code_details:via_upstream      connection_termination_details:-        upstream_transport_failure_reason:-     bytes_received:0        bytes_sent:135  duration:4      x-envoy-upstream-service-time:4 x-forwarded-for:-       user-agent:curl/7.73.0-DEV      x-request-id:84961386-6d84-929d-98bd-c5aee93b5c88       authority:httpbin:8000  upstream_host:10.44.1.27:80     upstream_cluster:outbound|8000||httpbin.foo.svc.cluster.local      upstream_local_address:10.44.1.23:37652 downstream_local_address:10.0.45.184:8000       downstream_remote_address:10.44.1.23:46520      requested_server_name:- route_name:default
```
