FROM golang:alpine AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/cache
COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache make git

WORKDIR /gin_demo
COPY . .
RUN go clean && GO111MODULE=on go bulid -ldflags "-s -w" -o docker/build/main cmd/main.go

COPY --from=builder /gin_demo/docker/build/main      /gin_demo/docker/build/

CMD ["/gin_demo/docker/build/main"]