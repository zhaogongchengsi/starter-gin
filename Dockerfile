
# ! Under development
FROM ubuntu
LABEL MAINTAINER="zzh1586169624@163.com"

FROM golang:alpine
WORKDIR /usr/src/app

#COPY ./web /usr/src/app/web
COPY . /usr/src/app/server
COPY ./config.yaml /usr/src/app/server/config.yaml

RUN echo "LANG=en_US.utf8" > /etc/locale.conf  \
    && go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
EXPOSE 80

# FROM redis
# FROM mysql