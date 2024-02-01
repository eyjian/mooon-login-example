# Copyright 2024, Tencent Inc.
# Author: yijian
# Date: 2024/02/02

all: mooon_login_example

mooon_login_example: mooon_login.go
	go build -o $@ $<

.PHONY: rpc tidy clean

clean:
	rm -fr mooon_login_example

rpc:
	goctl rpc protoc mooon_login.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=go_zero

tidy:
	go mod tidy
