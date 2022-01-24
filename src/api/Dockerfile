FROM golang:1.17.6

RUN apt-get update && \
    apt-get install -y locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9

WORKDIR /root/log

WORKDIR /root/api
COPY ./api /root/api/

RUN go mod init api && \
    go mod tidy

CMD ["/bin/sh", "/root/api/go.sh"]
