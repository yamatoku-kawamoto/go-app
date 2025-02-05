# Stage 1: Node.js (Vite) build
FROM node:20.17.0-alpine as vite-builder
WORKDIR /app

RUN chown -R node:node /usr/local
RUN chown -R node:node /app
USER node

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

# pnpm をインストール
RUN npm install -g pnpm

# 依存関係をインストール
COPY --chown=node:node views/pnpm-lock.yaml ./
COPY --chown=node:node views/package.json ./
RUN pnpm install

# ソースコードをコピーし、ビルドを実行
COPY --chown=node:node views/ .

RUN pnpm build

# Stage 2: Golang build
FROM golang:1.23.3 as golang-builder
WORKDIR /app

# 依存関係をインストール
COPY go.mod go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN go mod download

# Golang アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -tags production ./cmd

# Stage 3: Final runtime image
FROM alpine:3.21.2

# Timezone を設定
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN echo "Asia/Tokyo" > /etc/timezone
ENV TZ=Asia/Tokyo

WORKDIR /app

# # Golang 実行ファイルをコピー
COPY --from=golang-builder /app/main ./main
RUN chmod +x ./main

# # フロントエンドの静的ファイルをコピー
COPY --from=vite-builder /app/dist ./views

# # 必要なポートを公開
EXPOSE 8080

# # Golang 実行ファイルを実行
CMD ["./main"]