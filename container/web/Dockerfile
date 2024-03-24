FROM nginx:alpine

# 環境変数を設定
ENV TZ="Asia/Tokyo"

# 設定ファイルをコンテナ内にコピー
COPY /conf.d/default.conf /etc/nginx/default.conf
COPY /ssl/server.crt /etc/nginx/ssl/server.crt
COPY /ssl/server.csr /etc/nginx/ssl/server.csr
COPY /ssl/server.key /etc/nginx/ssl/server.key