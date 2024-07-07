FROM golang:1.22.5 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/www ./cmd

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/bin/www ./bin/www
COPY --from=builder /app/service_account_private_key.json .

EXPOSE 8080

CMD ["./bin/www"]
