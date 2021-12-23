FROM golang:1.15.2-alpine

# go moduleを使用
ENV GO111MODULE=on 

# アプリケーションを実行するディレクトリを指定
WORKDIR /go/src/github.com/ga-tcho/go-rest-api

# # 上記のディレクトリにgo.modとgo.sumをコピー
# COPY go.mod go.sum ./
# # 上記のファイルに変更がなければキャッシュ利用できる
# RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

# freshコマンドでサーバー起動
CMD ["fresh"]