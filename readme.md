


# Task Handler

Task Handler is a Go-based microservice designed to manage tasks. It provides a simple REST API for creating and retrieving tasks, using MongoDB for data storage.

## Technologies Used

- **Go**: The programming language used to build the service.
- **MongoDB**: The NoSQL database used for storing task data.
- **Docker**: Used for containerizing the application.
- **Kubernetes**: Optional, for orchestration and deployment.
- **cURL**: For testing the endpoints.

## Project Structure
```
/task-handler
│
├── cmd
│   └── main.go
│
├── internal
│   ├── config
│   │   └── config.go
│   ├── api
│   │   └── v1
│   │       └── task_handler.go
│   ├── repository
│   │   └── task_repository.go
│   ├── service
│   │   └── task_service.go
│   ├── pubsub
│   │   └── pubsub.go
│   └── server
│       └── server.go
│
├── pkg
│   ├── models
│   │   └── task.go
│   └── tools
│       └── utils.go
│
├── configs
│   └── config.yaml
│
├── scripts
│   └── migrate.sh
│
├── deployments
│   ├── Dockerfile
│   └── k8s
│
├── tests
│   ├── apitests
│   └── integration
│
├── go.mod
└── go.sum
```

## Local Setup

### Prerequisites

- Go 1.16+
- MongoDB (running locally or accessible remotely)
- Docker (optional, for containerization)
- Make sure to set the `MONGODB_URI` and `SERVER_PORT` environment variables if you want to override the default configuration.

### Steps

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/task-handler.git
    cd task-handler
    ```

2. **Set up the configuration**:
    - Update the `configs/config.yaml` file with your MongoDB URI and server port, or set the environment variables:
        ```yaml
        mongodb:
          uri: "mongodb://admin:password@localhost:27017"
        
        server:
          port: ":8080"
        ```

3. **Build the application**:
    ```sh
    go build -o task-handler ./cmd/main.go
    ```

4. **Run the application**:
    ```sh
    ./task-handler -env=development
    ```

### Docker Setup

1. **Build the Docker image**:
    ```sh
    docker build -t task-handler:latest .
    ```

2. **Run the Docker container**:
    ```sh
    docker run -d -p 8080:8080 -e MONGODB_URI="mongodb://admin:password@localhost:27017" --name task-handler task-handler:latest
    ```

### Kubernetes Setup

1. **Deploy to Kubernetes**:
    - Make sure your Kubernetes cluster is running.
    - Apply the Kubernetes manifests in the `deployments/k8s` directory:
        ```sh
        kubectl apply -f deployments/k8s
        ```

## Usage

### API Endpoints

- **GET /tasks**: Retrieve all tasks.
    ```sh
    curl http://localhost:8080/tasks
    ```

- **POST /task**: Create a new task.
    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"title": "New Task", "description": "Test the POST API", "completed": false}' http://localhost:8080/task
    ```

### Example Responses

- **GET /tasks**:
    ```json
    [
        {
            "id": "60c72b2f5f1b2c001c8e4b44",
            "title": "New Task",
            "description": "Test the POST API",
            "completed": false
        }
    ]
    ```

- **POST /task**:
    ```json
    {
        "InsertedID": "60c72b2f5f1b2c001c8e4b45"
    }
    ```

## Contributing

Contributions are welcome! Please create an issue or submit a pull request with your changes. Make sure to follow the project's coding standards and include tests for new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.