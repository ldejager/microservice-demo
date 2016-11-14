docker run -v "${PWD}/libs":/var/lib/drone -v /var/run/docker.sock:/var/run/docker.sock --env-file ./dronerc --restart=always -p 80:8000 -d --name=drone drone/drone:0.5
