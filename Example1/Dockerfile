FROM golang:1.23.0 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy && \
  CGO_ENABLED=0 GOOS=linux go build -o /api

FROM scratch

COPY --from=build /api /api

EXPOSE 3000

CMD ["/api"]