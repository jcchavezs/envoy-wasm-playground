.PHONY: build
build:
	tinygo build -o build/main.wasm -scheduler=none -target=wasi ./main.go