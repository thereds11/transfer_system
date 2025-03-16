FROM golang:1.21-alpine

# Install bash and postgres client for psql in entrypoint
RUN apk add --no-cache bash postgresql-client

# Set working directory
WORKDIR /app

# Copy everything into the container
COPY . .

# Install Go dependencies and build
RUN go mod tidy && go build -o main ./cmd/main.go

# Copy and fix permissions for entrypoint script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Set entrypoint
ENTRYPOINT ["/app/entrypoint.sh"]
