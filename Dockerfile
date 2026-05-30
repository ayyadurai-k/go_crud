# ---- Build stage ----
FROM golang:1.26-alpine AS builder

WORKDIR /src

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Static, stripped binaries. -trimpath removes local paths; -s -w drop debug info.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/server .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/migrate ./migrate

# ---- Runtime stage ----
FROM alpine:3.21

# ca-certificates: required to verify Supabase's TLS cert (sslmode=require).
# tzdata: for the Asia/Kolkata timezone. wget (busybox): used by the healthcheck.
RUN apk add --no-cache ca-certificates tzdata && \
    adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /out/server  /app/server
COPY --from=builder /out/migrate /app/migrate

USER appuser
EXPOSE 8000

ENTRYPOINT ["/app/server"]
