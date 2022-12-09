# istio-axslog

istio-axslog is parsed istio-proxy(envoy) access log and output in any format.

## Usage
```sh
$ istio-axslog --help
istio-axslog is parsed istio-proxy(envoy) access log and output in any format.

Usage:
  istio-axslog [flags]

Flags:
  -h, --help            help for istio-axslog
  -o, --output string   output format (default is json) supported formats are json, ltsv (default "json")
```

Output in json format (default):

```sh
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 3 1 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "127.0.0.1:80" inbound|8000|| 127.0.0.1:41854 10.44.1.27:80 10.44.1.23:37652 outbound_.8000_._.httpbin.foo.svc.cluster.local default' | istio-axslog
{"authority":"httpbin:8000","bytes_received":"0","bytes_sent":"135","connection_termination_details":"-","downstream_local_address":"10.44.1.27:80","downstream_remote_address":"10.44.1.23:37652","duration":"3","method":"GET","path":"/status/418","protocol":"HTTP/1.1","requested_server_name":"outbound_.8000_._.httpbin.foo.svc.cluster.local","response_code":"418","response_code_details":"via_upstream","response_flags":"-","route_name":"default","start_time":"2020-11-25T21:26:18.409Z","upstream_cluster":"inbound|8000||","upstream_host":"127.0.0.1:80","upstream_local_address":"127.0.0.1:41854","upstream_transport_failure_reason":"-","user-agent":"curl/7.73.0-DEV","x-envoy-upstream-service-time":"1","x-forwarded-for":"-","x-request-id":"84961386-6d84-929d-98bd-c5aee93b5c88"}
```

Output in ltsv format:
```sh
$ echo '[2020-11-25T21:26:18.409Z] "GET /status/418 HTTP/1.1" 418 - via_upstream - "-" 0 135 3 1 "-" "curl/7.73.0-DEV" "84961386-6d84-929d-98bd-c5aee93b5c88" "httpbin:8000" "127.0.0.1:80" inbound|8000|| 127.0.0.1:41854 10.44.1.27:80 10.44.1.23:37652 outbound_.8000_._.httpbin.foo.svc.cluster.local default' | istio-axlogs -o ltsv
authority:httpbin:8000  upstream_host:127.0.0.1:80      downstream_remote_address:10.44.1.23:37652      method:GET      response_code:418       response_code_details:via_upstream      bytes_received:0        bytes_sent:135  route_name:default      requested_server_name:outbound_.8000_._.httpbin.foo.svc.cluster.local      response_flags:-        x-forwarded-for:-       user-agent:curl/7.73.0-DEV      x-request-id:84961386-6d84-929d-98bd-c5aee93b5c88       upstream_cluster:inbound|8000|| start_time:2020-11-25T21:26:18.409Z     path:/status/418        upstream_transport_failure_reason:-        upstream_local_address:127.0.0.1:41854  downstream_local_address:10.44.1.27:80  protocol:HTTP/1.1       connection_termination_details:-        duration:3      x-envoy-upstream-service-time:1
```
