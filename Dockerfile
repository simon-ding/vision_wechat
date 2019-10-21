FROM golang:1.13-alpine AS build_base
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&\
    apk add gcc g++

WORKDIR /wechat/

ENV GO111MODULE=on GOPROXY=https://goproxy.cn

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -a -ldflags '-extldflags "-static"' -o ./cmd/server .

FROM alpine AS server
WORKDIR /wechat/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&\
    apk add tzdata &&\
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=build_base /wechat/cmd/server /bin/server

#配置文件
COPY ./config.yml /wechat/config.yml

ENTRYPOINT ["/bin/server"]