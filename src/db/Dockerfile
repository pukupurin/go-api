FROM mariadb:10.5

RUN apt-get update && \
    apt-get install -y locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

COPY ./my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf

COPY ./docker-entrypoint-initdb.d/ ./docker-entrypoint-initdb.d/
