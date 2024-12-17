# Event Tracking Microservice

This microservice allows users to log events and retrieve statistical insights about those events. It is built using **Golang** and stores data in **ClickHouse**.

## Features

- **Add Event**: Log user activity by saving events with a user ID, URL, and timestamp.
- **Get Statistics**: Retrieve unique user counts, total events, and the most visited URLs for a specific time period.

## Endpoints

### 1. Add Event

- **Endpoint**: `/v1/add_event`
- **Method**: `POST`
- **Request Body (JSON)**:

```json
{
  "user_id": "string",
  "url": "string",
  "timestamp": "string (ISO8601)"
}
```

- **Response**:
  - `200 OK`: Event successfully added.
  - `400 Bad Request`: Invalid request body.
  - `500 Internal Server Error`: Failed to save the event.

### 2. Get Statistics

- **Endpoint**: `/v1/get_stats`
- **Method**: `POST`
- **Query Parameters**:
  - `from` (required): Start of the period (ISO8601 format).
  - `to` (required): End of the period (ISO8601 format).

- **Response (JSON)**:

```json
{
  "unique_users": "int",
  "total_events": "int",
  "most_visited_urls": [
    {
      "url": "string",
      "count": "int"
    }
  ]
}
```

- **Errors**:
  - `400 Bad Request`: Missing or invalid query parameters.
  - `500 Internal Server Error`: Failed to retrieve statistics.


## Example Usage

### Adding an Event

```sh
curl -X POST http://localhost:8080/add_event \
-H "Content-Type: application/json" \
-d '{
  "user_id": "user123",
  "url": "https://example.com",
  "timestamp": "2024-12-16T10:00:00Z"
}'
```

### Retrieving Statistics

```sh
curl -X GET "http://localhost:8080/get_stats?from=2024-12-01T00:00:00Z&to=2024-12-15T23:59:59Z"
```

### Sample Response

```json
{
  "unique_users": 42,
  "total_events": 150,
  "most_visited_urls": [
    {
      "url": "https://example.com/page1",
      "count": 50
    },
    {
      "url": "https://example.com/page2",
      "count": 30
    }
  ]
}
```


## Setup and Run

### Prerequisites

- **Go** (v1.18 or later)
- **ClickHouse** (Running instance)
- **Docker** and **Docker Compose** (for database migrations)

### Steps

1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Start the services (with automatic migrations):
   ```sh
   docker-compose up --build
   ```

   The ClickHouse database will be set up automatically, and migrations will be applied.

4. Configure connection to ClickHouse in the code:
   Update the `setupClickHouse` function with the appropriate connection details if necessary.

5. Run the microservice:
   ```sh
   go run main.go
   ```

6. Access the endpoints at `http://localhost:8080`.

## Technologies Used

- **Golang**: For building the microservice.
- **ClickHouse**: For fast and efficient data storage and analytics.
- **Docker** (optional): To run ClickHouse as a container.

## Future Improvements

- Add authentication and authorization for endpoints.
- Add monitoring and logging for better observability.

---

**Author**: Egor Kutov