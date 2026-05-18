#############################
# Builder
#############################
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w" \
    -o api ./cmd/api

#############################
# Runtime
#############################
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates wget && \
    adduser -D -u 10001 appuser

USER 10001

COPY --from=builder --chown=10001:10001 /app/api .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./api"]
