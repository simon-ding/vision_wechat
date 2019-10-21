FROM golang:1.13-alpine AS build_base
RUN apk add gcc g++

WORKDIR /wechat/

ENV GO111MODULE=on GOPROXY=https://goproxy.io

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -a -ldflags '-extldflags "-static"' -o ./cmd/server .

FROM alpine AS server
WORKDIR /wechat/
RUN apk add tzdata &&\
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=build_base /wechat/cmd/server /bin/server

#配置文件
COPY ./config.yml /wechat/config.yml

ENTRYPOINT ["/bin/server"]