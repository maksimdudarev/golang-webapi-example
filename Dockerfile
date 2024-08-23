FROM golang:1.23.0

WORKDIR /usr/src/app

#RUN go mod init github.com/maksimdudarev/golang-webapi-example
#RUN go get github.com/gofiber/fiber/v2
#RUN go get gorm.io/gorm
#RUN go get gorm.io/driver/postgres

COPY . .
RUN go mod tidy