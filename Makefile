clean:
	@rm -rf build
	@rm -rf logs

.PHONY: build
build:
	mkdir -p build
	tinygo build -o build/main.wasm -scheduler=none -target=wasi ./main.go

run:
	docker-compose up --renew-anon-volumes