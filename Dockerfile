FROM golang:1.19.3-alpine3.16 AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3030

CMD [ "/app/main" ]