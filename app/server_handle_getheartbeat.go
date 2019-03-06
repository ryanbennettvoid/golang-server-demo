package main

import (
  "net/http"
)

func (server *Server) HandleGetHeartbeat(res http.ResponseWriter, req *http.Request) {
  response := Heartbeat{
    Message: "hello",
  }
  WriteJsonFromObject(response, true, res)
}
