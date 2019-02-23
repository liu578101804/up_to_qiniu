FROM alpine:latest

MAINTAINER liuy "578101804@qq.com"

# 解决 x509: certificate signed by unknown authority
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

WORKDIR /app

COPY ./build/app_linux /app/app

ENTRYPOINT ["/app/app"]
