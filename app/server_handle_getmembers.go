package main

import (
  "net/http"

  "github.com/ryanbennettvoid/golang-server-demo/repo"
)

func (server *Server) HandleGetMembers(res http.ResponseWriter, req *http.Request) {
  members, err := repo.GetMembers()
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }
  WriteJsonFromObject(members, true, res)
}
