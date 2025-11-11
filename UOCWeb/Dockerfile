# ---- build ----
FROM golang:1.25-alpine AS build
WORKDIR /app

# copia solo go.mod (go.sum puede no existir aún)
COPY go.mod ./
RUN go mod download

# ahora sí copia el resto del código (y generará go.sum si hace falta)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ---- runtime ----
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server /app/server
# si tienes estáticos:
# COPY --from=build /app/static /app/static
EXPOSE 8080
CMD ["/app/server"]
