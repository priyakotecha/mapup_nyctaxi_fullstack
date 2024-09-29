# mapup_nyctaxi_fullstack

# NYC Taxi Data Full Stack Application

This is a full-stack application that processes and analyzes the NYC Taxi dataset using a Go backend, a React frontend, Redis as a queue, and PostgreSQL for data storage.

## Features

- Backend built with Go.
- Frontend built with React.
- Redis for queue management and caching.
- PostgreSQL for data storage and analysis.
- Nginx for rate-limiting incoming requests.
- Role-based access (Admin, Manager, User).
- Dockerized for easy setup and deployment.

## Setup Instructions

### Prerequisites

- Docker
- Docker Compose
- Git

### 1. Clone the repository

git clone https://github.com/priyakotecha/mapup_nyctaxi_fullstack
cd fullstack-nyctaxi

### 2. Build and run docker container 

docker-compose up --build

This command will start all services including:

- Backend (Go) on port 8080
- Frontend (React) on port 3000
- PostgreSQL on port 5432
- Redis on port 6379
- Nginx (Rate Limiter) on port 80


### 4. Access the application
- Backend API: http://localhost:8080
- Frontend UI: http://localhost:3000
- Role-Based Endpoints
    /admin – Admin-level tasks, including user management and data modification.
    /manager – Access to analytics and viewing data.
    /user – Basic access to view insights and data summaries.