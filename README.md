## Logistics & Fleet Tracking Platform
This project is a high-impact, CQRS-based logistics system demonstrating real-time fleet routing, driver assignment, and route visibility across a distributed event-driven architecture.

## Tech Stack
- **CQRS + Event Sourcing**
- **Apache Kafka** for event streaming
- **MongoDB** for write-side persistence
- **Redis** for read-side projections
- **gRPC** for fast strongly typed APIs
- **Docker Compose** for local orchestration

## API Overview
### Command-Side (port 50051)
- CreateRoute(id, origin, destination)
- AssignDriver(route_id, driver_id)
- Publishes events to Kafka

### Query-Side (port 500052)
- GetRoute(id) - Read from Redis
- ListRoutes() - Query projected state

## Concepts Demonstrated
- CQRS: Command and Query Responsibilites Separated
- Event driven communication via Kafka
- Redis projections and view 
- Real-time fleet status tracking
- gRPC for performance + strict contracts

## Pending Tasks
- Mongo intergration
- Update assigned driver functions
- Update Route functionality

## Next Features
-  WebSocket updates for tracking
-  Prometheus + Grafana monitoring
-  REST gateway using gRPC-Gateway
