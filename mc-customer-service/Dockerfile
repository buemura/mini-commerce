# Base
FROM golang:1.26-alpine as base
WORKDIR /go/src/packages
COPY ./packages/ ./
RUN go mod download

WORKDIR /go/src/app
COPY ./mc-customer-service/go.mod ./
COPY ./mc-customer-service/go.sum ./
COPY ./mc-customer-service/ ./
RUN go mod download
RUN go build -o main ./cmd/grpc

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
