# This line specifies the base image for the Docker container
# - Uses the official Golang image version 1.24.0
# - Based on Alpine Linux which is a lightweight Linux distribution
# - Alpine variant helps keep the final image size small while providing necessary Go tools
FROM golang:1.24.0-alpine

WORKDIR /app

# Copy the entire backend directory
COPY . .

# Download dependencies
RUN go mod download

# Build the API service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api-service/main ./api-service

# Set the working directory to where the binary is
WORKDIR /app/api-service

EXPOSE 8080

CMD ["./main"]