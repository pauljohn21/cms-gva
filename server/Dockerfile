FROM golang:alpine as builder

WORKDIR /go/src/github.com/pauljohn21/cms-gva/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="pauljohn21@819820211@qq.com"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
WORKDIR /go/src/github.com/pauljohn21/cms-gva/server

COPY --from=0 /go/src/github.com/pauljohn21/cms-gva/server/server ./
COPY --from=0 /go/src/github.com/pauljohn21/cms-gva/server/resource ./resource/
COPY --from=0 /go/src/github.com/pauljohn21/cms-gva/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
