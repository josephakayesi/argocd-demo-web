# Stage 1: Build the Go binary
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd

# Stage 2: Slim runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/views ./views

EXPOSE 3000

CMD ["./server"]
