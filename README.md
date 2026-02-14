# Event Driven Commerce

## Overview

An ecommerce application built applying concepts of distributed systems, microservices, and event-driven architecture.

The **API Gateway** serves REST endpoints to the frontend and internally performs gRPC calls to backend microservices. Additionally, there is an asynchronous communication flow for order payment where `svc-order` publishes events to RabbitMQ and `svc-payment` processes them asynchronously, updating order status upon completion.

## Architecture Diagram

![Diagram](docs/event-driven-commerce.svg)

## Technologies

### Backend

- **Go** — all backend services
- **Echo** — HTTP framework (API Gateway)
- **gRPC / Protocol Buffers** — inter-service communication
- **RabbitMQ** — async messaging with dead-letter queues
- **PostgreSQL** — database per service
- **OpenTelemetry + Jaeger** — distributed tracing
- **Goose** — database migrations

### Frontend

- **TypeScript**
- **React 19** with **Vite**
- **Tailwind CSS**
- **Zustand** — state management
- **TanStack React Query** — server state / data fetching
- **React Router** — client-side routing

### Infrastructure

- **Docker** — multi-stage builds per service
- **Docker Compose** — local orchestration

## Services

| Service | Type | Port | Description |
|---|---|---|---|
| `api-gtw` | HTTP (REST) | 8080 | API Gateway — routes requests to backend services via gRPC |
| `svc-customer` | gRPC | 50051 | Customer registration, authentication (JWT + bcrypt) |
| `svc-product` | gRPC | 50050 | Product catalog and inventory management |
| `svc-order` | gRPC + Queue | 50052 | Order creation and lifecycle management |
| `svc-payment` | Queue Consumer | — | Processes payments asynchronously, publishes completion events |
| `web` | Vite Dev Server | 5173 | React frontend |

## Event-Driven Flow

The order payment flow uses RabbitMQ for asynchronous processing:

```
1. Frontend → POST /api/orders → API Gateway
2. API Gateway → gRPC → Order Service (creates order with WAITING_PAYMENT status)
3. Order Service → publishes event → RabbitMQ
4. Payment Service ← consumes event ← RabbitMQ
5. Payment Service processes payment and publishes to order.completed queue
6. Order Service ← consumes completion event ← updates order status
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/auth/signup` | Register a new customer |
| `POST` | `/api/auth/signin` | Sign in and receive JWT |
| `GET` | `/api/products` | List products (paginated) |
| `POST` | `/api/orders` | Create a new order (protected) |
| `GET` | `/api/orders` | Get user orders (protected) |

## How to Run

### Prerequisites

- Docker and Docker Compose
- Go 1.22+
- Node.js 18+

### Start Infrastructure

```bash
# Starts PostgreSQL, RabbitMQ, and Jaeger
sh scripts/env_up.sh
```

### Run Services

Each service can be run individually from its directory:

```bash
# Example: run the API Gateway
cd apps/api-gtw && make run

# Example: run the Order Service
cd apps/svc-order && make run
```

### Run Frontend

```bash
cd apps/web
npm install
npm run dev
```

### Useful URLs

| Service | URL |
|---|---|
| API Gateway | http://localhost:8080 |
| Jaeger UI | http://localhost:16686 |
| RabbitMQ Management | http://localhost:15672 |

## Screenshots

### Products Page

![Products](docs/fe-products.png)

### Checkout

![Checkout](docs/fe-checkout.png)

### Order List

![Order List](docs/fe-order-list.png)

### Order — Waiting Payment

![Order Waiting Payment](docs/fe-order-waiting-payment.png)

### Order — Completed

![Order Completed](docs/fe-order-completed.png)

## Project Structure

```
├── apps/
│   ├── api-gtw/          # REST API Gateway (Echo)
│   ├── svc-customer/     # Customer Service (gRPC)
│   ├── svc-order/        # Order Service (gRPC + Queue Consumer)
│   ├── svc-payment/      # Payment Service (Queue Consumer)
│   ├── svc-product/      # Product Service (gRPC)
│   └── web/              # React Frontend (Vite)
├── packages/
│   ├── grpc/             # Shared protobuf definitions
│   └── tracing/          # OpenTelemetry tracing utilities
├── configs/              # Configuration files
├── scripts/              # Deployment and setup scripts
└── docs/                 # Architecture diagrams and screenshots
```
