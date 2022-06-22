FROM golang:1.18.2-alpine3.16

# タイムゾーン設定
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /root/api

COPY ./src/go.* /root/api/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./src /root/api

CMD ["tail", "-f", "/dev/null"]
