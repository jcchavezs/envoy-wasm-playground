# Envoy WASM test

This project contains a setup to debug envoy request lifecycle for WASM filters.

## Getting started

```bash
make build run

curl -i localhost:8080
```

And you will get logs like

```bash
upstream_1   | 2023/01/04 10:47:43 GET /
upstream_1   |  > Accept: */*
upstream_1   |  > X-Forwarded-Proto: http
upstream_1   |  > X-Request-Id: 9554098a-6188-48cb-bb4d-6fae3ba02b13
upstream_1   |  > X-Envoy-Expected-Rq-Timeout-Ms: 15000
upstream_1   |  > User-Agent: curl/7.77.0
upstream_1   |
wasm-logs_1  | [2023-01-04 10:47:25.466382][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :authority: localhost:8080
wasm-logs_1  | [2023-01-04 10:47:25.466393][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :path: /
wasm-logs_1  | [2023-01-04 10:47:25.466394][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :method: GET
wasm-logs_1  | [2023-01-04 10:47:25.466395][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :scheme: http
wasm-logs_1  | [2023-01-04 10:47:25.466396][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> user-agent: curl/7.77.0
wasm-logs_1  | [2023-01-04 10:47:25.466397][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> accept: */*
wasm-logs_1  | [2023-01-04 10:47:25.466401][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> x-forwarded-proto: http
wasm-logs_1  | [2023-01-04 10:47:25.466402][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> x-request-id: 410127b7-570b-44ce-a353-857901f5fec9
wasm-logs_1  | [2023-01-04 10:47:25.496859][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response header <-- :status: 201
wasm-logs_1  | [2023-01-04 10:47:25.496868][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response header <-- content-type: application/json
wasm-logs_1  | [2023-01-04 10:47:25.496870][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response header <-- date: Wed, 04 Jan 2023 10:47:25 GMT
wasm-logs_1  | [2023-01-04 10:47:25.496871][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response header <-- content-length: 38
wasm-logs_1  | [2023-01-04 10:47:25.496872][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response header <-- x-envoy-upstream-service-time: 29
wasm-logs_1  | [2023-01-04 10:47:25.496940][27][debug][wasm] [source/extensions/common/wasm/context.cc:1167] wasm log test my_vm_id: about to get response body
wasm-logs_1  | [2023-01-04 10:47:25.496949][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: response body --> "{\n        \"message\": \"I am here\"\n    }"
wasm-logs_1  | [2023-01-04 10:47:25.497007][27][debug][wasm] [source/extensions/common/wasm/context.cc:1167] wasm log test my_vm_id: response body end of stream, body size: 0
wasm-logs_1  | [2023-01-04 10:47:43.139601][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :authority: localhost:8080
wasm-logs_1  | [2023-01-04 10:47:43.139628][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :path: /
wasm-logs_1  | [2023-01-04 10:47:43.139633][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :method: GET
wasm-logs_1  | [2023-01-04 10:47:43.139638][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> :scheme: http
wasm-logs_1  | [2023-01-04 10:47:43.139642][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> user-agent: curl/7.77.0
wasm-logs_1  | [2023-01-04 10:47:43.139646][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> accept: */*
wasm-logs_1  | [2023-01-04 10:47:43.139651][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> x-forwarded-proto: http
wasm-logs_1  | [2023-01-04 10:47:43.139656][27][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log test my_vm_id: request header --> x-request-id: 9554098a-6188-48cb-bb4d-6fae3ba02b13
```
