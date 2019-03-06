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

runServer()
{
  go run ./main.go
}

runTest()
{
  go test
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