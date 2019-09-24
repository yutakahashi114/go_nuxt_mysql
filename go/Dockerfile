# 開発
FROM golang:latest as dev
RUN apt-get update \
 && apt-get install -y vim \
 && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# 本番ステージング用
FROM golang:latest as builder
WORKDIR /go/src/app/
COPY ./src/app .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build server.go

# 本番
FROM alpine:latest as prod
EXPOSE 8080
WORKDIR /api
COPY --from=builder /go/src/app/server .
CMD ["./server"]
