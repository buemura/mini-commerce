# Base
FROM golang:1.26-alpine3.19 as base
WORKDIR /go/src/packages
COPY ./packages/ ./
RUN go mod download

WORKDIR /go/src/app
COPY ./mc-product-service/go.mod ./
COPY ./mc-product-service/go.sum ./
COPY ./mc-product-service/ ./
RUN go mod download
RUN go build -o main ./cmd/grpc

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
