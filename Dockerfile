FROM alpine:latest

MAINTAINER liuy "578101804@qq.com"

WORKDIR /app

COPY ./build/app_linux /app/app

ENTRYPOINT ["/app/app"]
