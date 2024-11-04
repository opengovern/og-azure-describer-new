.PHONY: build

local-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/og-azure-describer-new main.go

docker:
	docker build -t 435670955331.dkr.ecr.us-east-2.amazonaws.com/og-azure-describer-new:latest .
	docker push 435670955331.dkr.ecr.us-east-2.amazonaws.com/og-azure-describer-new:latest

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/og-azure-cli ./command/main.go
