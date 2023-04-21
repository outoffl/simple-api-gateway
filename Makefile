.PHONY: test

SERVICE := api-gateway
#MAIN := cmd/openapi/main.go
MAIN := cmd
GOBIN = go

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

go-mod-up:
	@echo "[`date +"%Y-%m-%d %H:%M:%S"`] " $(GOBIN)
	@echo "[`date +"%Y-%m-%d %H:%M:%S"`] 更新依赖"
	$(GOBIN) mod tidy

init-build:
	@echo "[`date +"%Y-%m-%d %H:%M:%S"`] 初始化构建目录"
	rm -rf build && mkdir build
	cp -Rfp conf ./build
	@echo "[`date +"%Y-%m-%d %H:%M:%S"`] 开始构建"

run:
	./build/bin/$(SERVICE)

build: go-mod-up init-build
	$(GOBIN) build -v -gcflags "-N" -o ./build/bin/$(SERVICE) ./$(MAIN)

test: build run

