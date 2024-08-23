FROM golang:1.23.0

WORKDIR /usr/src/app

RUN go mod init github.com/maksimdudarev/golang-webapi-example
RUN go get github.com/gofiber/fiber/v2

COPY . .
RUN go mod tidy