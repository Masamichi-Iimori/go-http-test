FROM golang:1.13

RUN apt-get update -qq && \
  apt-get install -y mariadb-client vim

ENV APP_ROOT /tmp/mnt/src
ENV DOCKER_ROOT /tmp/mnt/docker

## 作業ディレクトリ
WORKDIR $APP_ROOT

# モジュール管理のファイルをコピー
COPY ./src/go.mod $APP_ROOT
COPY ./src/go.sum $APP_ROOT

# 外部パッケージのダウンロード
RUN go mod download all

EXPOSE 9000

# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN go build -o main src/main.go

# ENTRYPOINT ["/app/main"]
# EXPOSE 80