# Build go
FROM golang:1.25.0-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN GOEXPERIMENT=jsonv2 go mod download
RUN GOEXPERIMENT=jsonv2 go build -v -o NeXT-Server -tags "xray"

# Release
FROM  alpine
# 安装必要的工具包
RUN  apk --update --no-cache add tzdata ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN mkdir /etc/NeXT-Server/
COPY --from=builder /app/NeXT-Server /usr/local/bin

ENTRYPOINT [ "NeXT-Server", "server", "--config", "/etc/NeXT-Server/config.json"]
