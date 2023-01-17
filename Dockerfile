# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder
WORKDIR /source
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./executable

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -hS peon -G appgroup --disabled-password
USER peon
WORKDIR /home/peon
COPY --chown=peon:peon --from=builder /source/executable .
ENV BROKER_DBHOST=host.docker.internal
ENV APP_PROXIES=host.docker.internal
ENV CACHE_HOST=host.docker.internal
ENV COOKIE_DOMAIN=host.docker.internal
EXPOSE 8080
ENTRYPOINT ["./executable"]
