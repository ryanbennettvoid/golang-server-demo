package main

import (
  "net/http"

  "github.com/go-chi/chi"
  "github.com/ryanbennettvoid/golang-server-demo/repo"
)

func (server *Server) HandleGetMemberById(res http.ResponseWriter, req *http.Request) {
  memberId := chi.URLParam(req, "id")
  member, err := repo.GetMemberById(memberId)
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }
  WriteJsonFromObject(member, true, res)
}
