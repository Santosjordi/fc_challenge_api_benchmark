# Step 1: build stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

# build a statically linked binary
RUN CGO_ENABLED=0 go build -o loadtest ./cmd/benchmark

# Step 2: final stage (tiny image)
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/loadtest .

# This lets docker run pass flags directly
ENTRYPOINT ["./loadtest"]
