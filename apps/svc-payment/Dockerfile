# Base
FROM golang:1.26-alpine3.19 as base

WORKDIR /go/src/app
COPY ./apps/svc-payment/go.mod ./
COPY ./apps/svc-payment/go.sum ./
COPY ./apps/svc-payment/ ./
RUN go mod download
RUN go build -o main ./cmd/queue

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
