# Use official Golang image with Go 1.25.1 as builder
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task_vantra .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/task_vantra .

EXPOSE 8080

CMD ["./task_vantra"]
