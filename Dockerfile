FROM golang:1.20.3-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o server

FROM alpine:3.16
COPY --from=builder /app/server /server
CMD ["/server"]
