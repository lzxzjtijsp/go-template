# syntax=docker/dockerfile:1

# ベースイメージとして公式のGoイメージを使用
FROM golang:1.20.4-alpine as builder

# 作業ディレクトリを設定
WORKDIR /go-template

# 必要なパッケージをインストール
RUN apk add --no-cache git

# ソースコードをコピー
COPY . .

# Goの依存関係をダウンロード
RUN go mod download
RUN go mod tidy
RUN go mod verify

# ビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 新しいステージを作成し、実行可能ファイルをコピー
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go-template/main .

# Application
ENV APP_PORT="8080"
# Database credentials
ENV DB_DRIVER="mysql"
ENV DB_USER="root"
ENV DB_PASSWORD=""
ENV DB_HOST="127.0.0.1"
ENV DB_PORT="3306"
ENV DB_NAME="go-temlate"
ENV DB_CHARSET="utf8"
ENV DB_MAX_OPEN_CONNS="2"

# アプリケーションを実行
CMD ["./main"]
