# API

The purpose of the demo is to take a service and break it down into smaller pieces to allow for scalability by decoupling the API in this example from the database.

Using Gin as a web framework the goal is to create a simple task API which has endpoints for creating and listing tasks and getting a basic health status back.

## Development

### Building

You have two options in building the binary, the first being locally which uses your host operation systems golang path, libraries etc., the other being to use the pre-built binary in the docker container.

#### Locally
```
go build
```

#### Docker

```
docker build -t api .
```

### Running

Once the binary has been built, run it like so locally. Alternatively, run the docker container you have created in the step above.

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
curl -v -H "Content-Type: application/json" -X POST -d '{"Name": "", "Description": ""}' localhost:8000/tasks
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

- Create sample tests which is triggered by the build toolset
- Create Kubernetes resource controller and service etc. definitions
