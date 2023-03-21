FROM golang:alpine as StarterServer
LABEL MAINTAINER="zzg<zzh1586169624@163.com>"
WORKDIR /usr/src/app
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .
ENTRYPOINT ./server -c ./ -t yaml -n config.pro
EXPOSE 80

#ENTRYPOINT ./server -c ./ -t yaml -n config.pro
#FROM alpine:latest
#WORKDIR /usr/src/app
#COPY --from=0 /usr/src/app ./
#COPY --from=0 /usr/src/app/resource ./resource/
#COPY --from=0 /usr/src/app/config.pro.yaml ./


