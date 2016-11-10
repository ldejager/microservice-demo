# API

The purpose of the demo is to take a service and break it down into smaller pieces to allow for scalability by decoupling the API in this example from the database.

Using Gin as a web framework the goal is to create a simple task API which has endpoints for creating and listing tasks and getting a basic health status back.

## Usage

### Building

```
go build
```

### Running

Once the binary has been built, run it like so;

```
./api
```

### Configuration

Currently there is no configuration for the API, however, once the TODO items have been done I'd expect database and other configuration values to be passed in as environment variables.

### Interacting

#### Microservices Welcome
Simple, generic welcome page
```
curl -v localhost:8000/
```

#### List Tasks
List all tasks...
```
curl -v localhost:8000/tasks
```

#### Create Task
Create a simple task on the API via an HTTP POST (JSON)
```
curl -v -X POST localhost:8000/ -d '{}'
```

#### View Task Details
View specific task details, takes an task ID as argument
```
curl -v localhost:8000/tasks/<id>
```

#### Health Stats
End to end testing of internal components, providing a summary of overall health.
```
curl -v localhost:8000/health
```

#### Application Monitoring
Useful to monitor application availability from external monitoring systems, i.e. Kubernetes livenessProbe.
```
curl -v localhost:8000/ping
```

For example;

```
livenessProbe:
  httpGet:
    path: /ping
    port: 8000
  initialDelaySeconds: 30
  timeoutSeconds: 1
```

## TODO

- Fix mapping bug
- Abstract sqlite database from the service
- Create a real health check
- Create build toolset for the API
- Create sample tests which is trigger by the build toolset
- Dockerise the application (in progress)
- Create Kubernetes resource controller and service etc. definitions
