FROM golang:1.25.5-alpine AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/. .

RUN CGO_ENABLED=0 go build -o main backend/main.go

FROM scratch
WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080
CMD ["./main"]
