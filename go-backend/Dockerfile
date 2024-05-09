# Build stage
FROM golang:1.22.0-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/db/migrations ./db/migrations

EXPOSE 8080
CMD ["/app/main"]