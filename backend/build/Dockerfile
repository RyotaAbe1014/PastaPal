# GoのAlpineベースイメージ
FROM golang:1.23-alpine3.20

# パッケージ情報の更新とgitのインストール
RUN apk update && apk add --no-cache git

# タイムゾーンの設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# 作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコピーして依存関係を解決
# COPY go.mod go.sum ./
RUN go mod init backend
RUN go mod tidy

# AirのバイナリをインストールしてPATHに追加
RUN go install github.com/air-verse/air@latest && mv /go/bin/air /usr/local/bin/air

# sqlcのバイナリをインストールしてPATHに追加
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && mv /go/bin/sqlc /usr/local/bin/sqlc

EXPOSE 8000

# コンテナ起動時のデフォルトコマンド
CMD ["air", "-c", "build/.air.toml"]
