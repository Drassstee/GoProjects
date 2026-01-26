FROM golang:1.25-alpine AS builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o root ./cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /service/root ./root

EXPOSE 8080

CMD ["./root"]