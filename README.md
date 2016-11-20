[![Build Status](https://droneci.rwxlabs.io/api/badges/ldejager/microservice-demo/status.svg)](https://droneci.rwxlabs.io/ldejager/microservice-demo)  [![](https://images.microbadger.com/badges/image/ldejager/microservice-demo:d305948e.svg)](https://microbadger.com/images/ldejager/microservice-demo:d305948e "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/ldejager/microservice-demo.svg)](https://microbadger.com/images/ldejager/microservice-demo "Get your own version badge on microbadger.com")

# API

The purpose of the demo is to take a service and break it down into smaller pieces to allow for scalability by decoupling the API in this example from the database.

Using Gin as a web framework the goal is to create a simple task API which has endpoints for creating and listing tasks and getting a basic health status back.

## Development

### Building

The repo provides a Makefile which can use used to build cross platform binaries of the microservice-demo binary. From within the repository, run the following;

```
make
```

### Configuration

The API expects the following environment variables to be passed in before it will start;

`DB_USERNAME`
`DB_PASSWORD`
`DB_HOSTNAME`
`DB_DATABASE`

You will need to have this setup prior to launching the API.

### Running

Once the binary has been built, you can either run it locally or via a docker container.

For example;

```
export DB_USERNAME=test
export DB_PASSWORD=test
export DB_HOSTNAME=test
export DB_DATABASE=test
release/linux/amd64/microservice-demo
```

Running the docker container is just as simple;

```
docker run -d -p80:8000 -e DB_USERNAME=test -e DB_PASSWORD=test -e DB_HOSTNAME=test -e DB_DATABASE=test ldejager/microservices-demo
```

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
