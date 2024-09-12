FROM golang:1.23.0-alpine AS build

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

RUN go build -o /api

FROM gcr.io/distroless/static AS final

COPY --from=build /api /api

EXPOSE 3000

ENTRYPOINT ["/api"]