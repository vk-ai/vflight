# V-Flight

V-Flight is a microservices-based flight management system built using Go. The system consists of multiple services that work together to provide a comprehensive flight management solution.

## Project Structure

```
.
├── backend/
│   ├── api-service/    # Main API service
│   ├── user-service/   # User management service
│   ├── commons/        # Shared utilities and components
│   └── vendor/         # Dependencies
├── docker-compose.yml  # Docker services configuration
└── .env               # Environment configuration
```

## Services

1. **API Service** (Port 8080)
   - Main API service handling flight-related operations
   - Connects to PostgreSQL database

2. **User Service** (Port 8081)
   - Handles user management and authentication
   - Connects to PostgreSQL database

3. **PostgreSQL Database** (Port 5432)
   - Stores application data
   - Uses persistent volume for data storage

## Prerequisites

- Docker and Docker Compose
- Go 1.x (for local development)
- Make (optional, for using Makefile commands)

## Environment Configuration

The project uses the following environment variables (configured in `.env`):

```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=vflight
API_SERVICE_PORT=8080
USER_SERVICE_PORT=8081
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd v-flight
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up -d
   ```

   This will:
   - Build and start all services
   - Create necessary networks and volumes
   - Initialize the PostgreSQL database

3. Verify services are running:
   ```bash
   docker-compose ps
   ```

## Service Endpoints

1. API Service (http://localhost:8080)
   - Main API endpoints for flight management

2. User Service (http://localhost:8081)
   - User management and authentication endpoints

## Development

For local development:

1. Install dependencies:
   ```bash
   cd backend
   go mod download
   ```

2. Run services individually:
   ```bash
   # Run API service
   cd backend/api-service
   go run main.go

   # Run User service
   cd backend/user-service
   go run main.go
   ```

## Stopping the Services

To stop all services:
```bash
docker-compose down
```

To stop services and remove volumes:
```bash
docker-compose down -v
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

[Add your license information here]
