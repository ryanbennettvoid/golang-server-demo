#!/bin/bash

CMD=$1

usage()
{
  echo "USAGE: $0 <command>"
  echo "valid commands are:"
  echo "  server"
  echo "  test"
  echo "  docker"
}

noteMongo()
{
  echo "NOTE: make sure that mongo is running (mongo.sh)"
}

runServer()
{
  noteMongo
  go run ./app
}

runTest()
{
  noteMongo
  go test -failfast -v ./...
}

runDocker()
{
  docker-compose -f ./docker/docker-compose.yml up
}

case "$CMD" in
  "server") runServer;;
  "test") runTest;;
  "docker") runDocker;;
  *)
    echo "invalid command: $1"
    usage
    ;;
esac