# Golang Server

Run server: `./run server`
Run test: `./run test`
Run server in docker: `./run docker`

NOTES:
- to install dependencies on local machine, run `go get ./...`
- running on local machine requires mongo to be running (`./mongo.sh`)
- ensure that only a single mongo instance is running so ports don't clash (remove all containers with `docker rm -f $(docker ps -qa)`)