FROM golang:1.23.1

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src/app
RUN go mod init chat
RUN go mod tidy
RUN echo "deb http://archive.debian.org/debian/ stretch main" > /etc/apt/sources.list \
    && echo "deb http://archive.debian.org/debian-security stretch/updates main" >> /etc/apt/sources.list \
    && apt-get update
# RUN go get -u github.com/go-sql-driver/mysql
RUN go get "github.com/lib/pq"
RUN go install github.com/beego/bee/v2@latest
RUN go get "github.com/beego/beego/"
# RUN go install github.com/beego/beego/v2@latest

# 自動的にYesでデータベースをマイグレーション
# RUN echo "Yes" | bee migrate
# beeコマンドがインストールされているか確認
RUN which bee