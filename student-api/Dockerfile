# Stage 1: Builder
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install CA certs and create non-root user
RUN apk add --no-cache ca-certificates && \
    addgroup -S student && \
    adduser -S -D -G student -s /bin/sh -u 1000 student

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build statically-linked binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -a -installsuffix cgo \
    -o student-api ./cmd/student-api/main.go

# Stage 2: Final image
FROM scratch AS final

WORKDIR /app

# Copy only what's needed
COPY --from=builder /app/student-api .
COPY --from=builder /app/config ./config
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Use non-root user
USER 1000

EXPOSE 8082

# Pass config path as argument
ENTRYPOINT ["./student-api", "--config=config/dev.yaml"]
