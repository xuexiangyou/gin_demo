FROM golang:alpine AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/cache
COPY go.mod go.sum ./
RUN go mod download

# 切换apk为阿里云的源 腾讯服务器太慢了
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache make git

WORKDIR /gin_demo
COPY . .
RUN go clean && GO111MODULE=on go build -ldflags "-s -w" -o docker/build/main cmd/main.go

FROM alpine as gin_demo
WORKDIR /gin_demo
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /gin_demo/docker/build/main      /gin_demo/docker/build/

CMD ["/gin_demo/docker/build/main"]