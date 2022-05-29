FROM golang:1.18.2-alpine3.16

# タイムゾーン設定
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /root/api

COPY ./src/go.* /root/api/
RUN go mod download

COPY ./src /root/api

CMD ["/bin/sh", "/root/api/go.sh"]
