FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/public ./public

EXPOSE 8080
CMD ["./server"]