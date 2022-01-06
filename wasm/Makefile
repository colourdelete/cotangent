all: wasm_exec.js cotangent.wasm server;

wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm_exec.js

cotangent.wasm:
	GOOS=js GOARCH=wasm go build -o cotangent.wasm ../main.go

server:
	go run server.go

.PHONY: server
