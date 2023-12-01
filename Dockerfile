FROM golang:1.21-alpine as builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV CGO_ENABLED 0
RUN go build -o /assets/in ./cmd/in
RUN go build -o /assets/out ./cmd/out
RUN go build -o /assets/check ./cmd/check

FROM alpine:3.18 as resource
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

USER 1001:1001
