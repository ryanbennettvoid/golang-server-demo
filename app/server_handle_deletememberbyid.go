package main

import (
  "net/http"

  "github.com/go-chi/chi"
  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo"
)

func (server *Server) HandleDeleteMemberById(res http.ResponseWriter, req *http.Request) {

  memberId := chi.URLParam(req, "id")

  err := repo.DeleteMemberById(memberId)
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }

  WriteJsonFromObject(SuccessResponse{
    Message: "member updated",
  }, true, res)

}
