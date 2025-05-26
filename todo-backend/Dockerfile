# Build stage
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080

CMD ["./server"]