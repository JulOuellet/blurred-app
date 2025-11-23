# --------------------------------------------------------
# STAGE 1: Build (Compiles CSS, Templ, and Go Binary)
# --------------------------------------------------------
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache curl git libc6-compat

WORKDIR /app

# 1. Install Templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# 2. Install Tailwind CSS Standalone CLI
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# 3. Download Go Dependencies (Caching layer)
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy Source Code
COPY . .

# 5. Generate Templ files
RUN templ generate

# 6. Generate Tailwind CSS
RUN tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css -c tailwind.config.js --minify

# 7. Build Go Binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --------------------------------------------------------
# STAGE 2: Runtime (The actual production container)
# --------------------------------------------------------
FROM alpine:latest

# Install certificates (needed for HTTPS/SSL DB connections)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the migrations folder so the Go app can find them
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations

# Copy static assets (CSS, images, JS)
COPY --from=builder /app/assets ./assets/

# Expose the port (for documentation, Railway overrides this)
EXPOSE 8080

# Run the binary
CMD ["./main"]
