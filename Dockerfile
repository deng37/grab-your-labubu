# Stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Build binary statis
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /root/
# Copy binary dari stage builder
COPY --from=builder /app/main .
# Copy folder assets dan index.html agar bisa dilayani
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/index.html .

EXPOSE 8080
CMD ["./main"]