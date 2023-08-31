# goバージョン
FROM --platform=linux/arm64 golang:1.21.0-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# full-moon-backendディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# gqlgenのインストール
RUN go install github.com/99designs/gqlgen@latest
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports
# golang-migrateのインストール
RUN go install -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# shell scriptの実行
COPY /docker/app/ /go/src/app/
RUN chmod +x ./docker/app/entrypoint.sh
ENTRYPOINT [ "./docker/app/entrypoint.sh" ]

