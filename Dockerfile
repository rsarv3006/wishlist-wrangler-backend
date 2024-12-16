FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

# RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

# RUN swag init -d ./,../../shared --parseDependency true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./services/user-service

FROM alpine:latest

WORKDIR /root/

COPY ./services/user-service/config.yml .

COPY ./services/user-service/docs/swagger.json .

COPY --from=builder /app/user-service .

EXPOSE 8080

CMD ["./user-service"]
