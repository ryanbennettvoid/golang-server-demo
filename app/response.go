package main

const ERROR_MISSING_REQUEST_BODY = "missing request body"

type ErrorResponse struct {
  Message string `json:"message"`
}

type SuccessResponse struct {
  Message string `json:"message"`
}
