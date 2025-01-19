FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .env.docker /app/.env

RUN apk add --no-cache bash postgresql-client

EXPOSE 7888

RUN go build -o main cmd/country/main.go

CMD ["./main"]