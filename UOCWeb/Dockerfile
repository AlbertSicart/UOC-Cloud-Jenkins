# Etapa de build
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app .

# Etapa de ejecución
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
COPY static ./static
EXPOSE 8080
CMD ["./app"]
