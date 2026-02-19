FROM golang:1.25.5-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine:latest AS prod
WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080
CMD ["./main"]
