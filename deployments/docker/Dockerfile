FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nextpress-backend ./cmd/nextpress

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/nextpress-backend /app/nextpress-backend

ENV APP_ENV=prod

EXPOSE 8080

ENTRYPOINT ["/app/nextpress-backend"]

