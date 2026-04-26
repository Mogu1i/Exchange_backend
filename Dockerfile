# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/exchangeapp .

FROM alpine:3.20
ARG APP_VERSION=dev
LABEL org.opencontainers.image.title="exchangeapp"
LABEL org.opencontainers.image.version=$APP_VERSION

RUN adduser -D -g '' app
WORKDIR /app

COPY --from=builder /out/exchangeapp /app/exchangeapp
COPY docker-entrypoint.sh /app/docker-entrypoint.sh

RUN chmod +x /app/docker-entrypoint.sh \
    && chown -R app:app /app

USER app
EXPOSE 3000

ENV CONFIG_PATH=/app/config/config.yml
ENV APP_VERSION=$APP_VERSION

ENTRYPOINT ["/app/docker-entrypoint.sh"]
