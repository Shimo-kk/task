FROM golang:1.19.2-alpine

# 環境変数を設定
ENV GO111MODULE on
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ワークディレクトリを指定
WORKDIR /src

# パッケージのインストール
RUN apk update && \
    apk --no-cache add git gcc musl-dev && \
    apk add bash && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ソースコードをコンテナにコピー
COPY /src /src

# 解放するポートを指定
EXPOSE 8000

#CMD [ "go", "run", "app/main.go" ]