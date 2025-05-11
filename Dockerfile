# Stage 1: Build the binaries
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build statically-linked binaries (no glibc dependency)
ENV CGO_ENABLED=0
RUN go build -o /bin/app ./cmd/app
RUN go build -o /bin/migrate ./cmd/database/migrations/main.go
RUN go build -o /bin/mock ./cmd/mock/main.go

# Stage 2: Minimal runtime image
FROM scratch

# Copy statically-built binaries from builder
COPY --from=builder /bin/app /bin/app
COPY --from=builder /bin/migrate /bin/migrate
COPY --from=builder /bin/mock /bin/mock

# Copy .env file (optional)
COPY .env /.env

# Default entrypoint
ENTRYPOINT ["/bin/app"]
