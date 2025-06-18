FROM golang:1.24-alpine AS builder
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest
COPY --from=builder /go/bin/goose /usr/bin/goose
ENTRYPOINT ["goose"]
