package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"

  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo"
)

func (server *Server) HandleCreateMember(res http.ResponseWriter, req *http.Request) {

  reqData, err := ioutil.ReadAll(req.Body)
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }

  if len(reqData) == 0 {
    WriteJsonFromObject(ErrorResponse{
      Message: ERROR_MISSING_REQUEST_BODY,
    }, false, res)
    return
  }

  var member repo.Member
  err = json.Unmarshal(reqData, &member)
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }

  if err := member.Validate(); err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }

  err = repo.InsertMember(member)
  if err != nil {
    WriteJsonFromObject(ErrorResponse{
      Message: err.Error(),
    }, false, res)
    return
  }

  WriteJsonFromObject(SuccessResponse{
    Message: "member created",
  }, true, res)
}
