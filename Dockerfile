
# ! Under development
FROM ubuntu
LABEL MAINTAINER="zzh1586169624@163.com"

FROM golang:alpine
WORKDIR /usr/src/app

RUN echo world

# FROM redis
# FROM mysql