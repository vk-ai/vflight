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

# Build the user service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/user-service/main ./user-service

# Set the working directory to where the binary is
WORKDIR /app/user-service

EXPOSE 8081

CMD ["./main"]