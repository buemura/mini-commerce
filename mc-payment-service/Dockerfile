# Base
FROM golang:1.26-alpine3.19 as base
WORKDIR /go/src/packages
COPY ./packages/ ./
RUN go mod download

WORKDIR /go/src/app
COPY ./mc-payment-service/go.mod ./
COPY ./mc-payment-service/go.sum ./
COPY ./mc-payment-service/ ./
RUN go mod download
RUN go build -o main ./cmd/queue

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
