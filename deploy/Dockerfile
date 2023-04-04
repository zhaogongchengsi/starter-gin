FROM golang:alpine as starter-server
LABEL MAINTAINER="zzg<zzh1586169624@163.com>"
WORKDIR /usr/src/app
COPY .. .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server . \
    && sudo mkdir "logs" \
    && sudo mkdir "upload"
ENTRYPOINT ./server -c ./ -t yaml -n config.local
EXPOSE 80
