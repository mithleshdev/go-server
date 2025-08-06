FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
