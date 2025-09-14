# ---------- Build Stage ----------
FROM --platform=linux/amd64 golang:1.23 AS builder

ENV GIT_TERMINAL_PROMPT=1 GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working dir
WORKDIR /app

# Copy go.mod dan go.sum dulu biar cache build lebih efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source
COPY . .

# Build binary
RUN go build -ldflags="-w -s" -o main ./cmd/main.go


# ---------- Runtime Stage ----------
FROM alpine:3.13.1

# Timezone
ENV TZ=Asia/Jakarta
RUN apk add --no-cache tzdata ca-certificates \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/main .

# Copy .env ke dalam image
COPY .env .env

EXPOSE 8080

ENTRYPOINT ["./main"]
