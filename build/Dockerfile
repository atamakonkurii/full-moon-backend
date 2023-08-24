# goバージョン
FROM golang:1.21.0-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# full-moon-backendディレクトリの作成
# RUN mkdir /go/src/github.com/full-moon-backend
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com
# gqlgenのインストール
RUN go install github.com/99designs/gqlgen@latest
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports