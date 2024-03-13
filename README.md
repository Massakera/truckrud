# Vehicle and Driver Management System

Welcome to Truckrud! This project provides a CRUD (Create, Read, Update, Delete) functionality for managing drivers and vehicles, along with a feature to link drivers with vehicles.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed the latest version of [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

### Installing and Running

To get the application up and running, follow these steps:

1. Clone the repository to your local machine
2. Navigate to the project directory:

```bash
cd truckrud
```

 3. Build and run the containers using Docker Compose:

 ```bash
 docker compose up --build
 ```

4. Once the containers are up and running, the API server should be accessible via `http://localhost:8080`.

### Usage

After running the containers, you can use the API endpoints to manage drivers and vehicles. The endpoints will allow you to:

- Add, retrieve, update, and delete driver records.
- Add, retrieve, update, and delete vehicle records.
- Link and unlink drivers with vehicles.

#### Tests

To run the tests, simply go to the tests folder and run `go test`.

the swagger doc is avaliable here: `http://localhost:8080/swagger/index.html` (access while running the project).

