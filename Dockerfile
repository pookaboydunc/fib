# --- Stage 1: Build ---
FROM golang:1.24.4 AS builder

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go binary
RUN make clean
RUN make build

# --- Stage 2: Run ---
FROM gcr.io/distroless/base-debian12

WORKDIR /root/

# Copy only the binary from builder
COPY --from=builder /app/bin/fib .

# Run service on container start
CMD ["./fib"]
