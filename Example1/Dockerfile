FROM golang:1.23.0

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

EXPOSE 3000

CMD ["/api"]