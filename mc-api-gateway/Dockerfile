# Base
FROM golang:1.26-alpine as base
WORKDIR /go/src/packages
COPY ./packages/ ./
RUN go mod download

WORKDIR /go/src/app
COPY ./mc-api-gateway/go.mod ./
COPY ./mc-api-gateway/go.sum ./
COPY ./mc-api-gateway/ ./
RUN go mod download
RUN go build -o main ./cmd/http

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
