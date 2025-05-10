# Go-Caching

A flexible multi-level caching implementation in Go, demonstrating different caching strategies using in-memory cache and Redis.

## Features

- **L1 Cache**: In-memory caching using [patrickmn/go-cache](https://github.com/patrickmn/go-cache)
- **L2 Cache**: Distributed caching using Redis
- **L1-L2 Cache**: Multi-level caching combining both in-memory and Redis
- RESTful API interface using Gin framework
- Configurable cache expiration times
- Simulated database delay for demonstration

## Requirements

- Go 1.23.1 or higher
- Redis server (for L2 and L1-L2 implementations)
- Dependencies (automatically installed via go mod):
  - github.com/gin-gonic/gin
  - github.com/patrickmn/go-cache
  - github.com/redis/go-redis/v9

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-caching.git
cd go-caching
```

2. Install dependencies:
```bash
go mod download
```

3. Ensure Redis is running (for L2 and L1-L2 implementations):
```bash
redis-server
```

## Usage

The project provides three different caching implementations:

### 1. L1 (In-Memory) Cache

```go
// Run the L1 cache server
go run l1-go-cache.go
```

### 2. L2 (Redis) Cache

```go
// Run the L2 cache server
go run l2-go-cache.go
```

### 3. L1-L2 (Multi-Level) Cache

```go
// Run the L1-L2 cache server
go run l1-l2-go-cache.go
```

## API Endpoints

All implementations expose the following endpoint:

### GET /data/:key

Retrieves data for the specified key.

- **URL Parameter**: key (string)
- **Response Format**:
```json
{
    "message": "ok",
    "data": "cached_value"
}
```

## Caching Strategies

1. **L1 Cache (In-Memory)**:
   - Uses in-memory cache with 1-minute default expiration
   - 2-minute cleanup interval
   - Ideal for frequently accessed data

2. **L2 Cache (Redis)**:
   - Uses Redis with 5-minute expiration
   - Suitable for distributed systems
   - Persistent across application restarts

3. **L1-L2 Cache (Multi-Level)**:
   - Combines both in-memory and Redis caching
   - Checks L1 cache first, then L2
   - Updates both levels on cache miss
   - Provides optimal balance between performance and reliability

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
