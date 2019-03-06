package main

import (
  "encoding/json"
  "net/http"
)

func WriteJsonFromBytes(data []byte, success bool, res http.ResponseWriter) {
  res.Header().Set("Content-Type", "application/json")
  if success {
    res.WriteHeader(http.StatusOK)
  } else {
    res.WriteHeader(http.StatusInternalServerError)
  }
  res.Write(data)
}

func WriteJsonFromObject(obj interface{}, success bool, res http.ResponseWriter) {
  res.Header().Set("Content-Type", "application/json")
  if success {
    res.WriteHeader(http.StatusOK)
  } else {
    res.WriteHeader(http.StatusInternalServerError)
  }
  data, err := json.Marshal(obj)
  if err != nil {
    panic(err)
    return
  }
  res.Write(data)
}
