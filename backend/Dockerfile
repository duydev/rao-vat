FROM golang:1.22.5 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o www ./cmd

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/www .

EXPOSE 8080

CMD ["./www"]
