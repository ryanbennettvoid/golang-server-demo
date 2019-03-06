package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Client struct {
  BaseUrl string
}

func NewClient() *Client {
  return &Client{
    BaseUrl: fmt.Sprintf("http://localhost:%d", DEFAULT_PORT),
  }
}

type ClientRequestOptions struct {
  Method   string
  Endpoint string
  Headers  map[string]string
  Body     map[string]string
}

func NewClientRequestOptions() ClientRequestOptions {
  return ClientRequestOptions{
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
    Body: map[string]string{},
  }
}

func (c *Client) MakeRequest(options ClientRequestOptions) (string, error) {
  url := c.BaseUrl + options.Endpoint
  requestBody, err := json.Marshal(options.Body)
  resp, err := http.NewRequest(options.Method, url, bytes.NewBuffer(requestBody))
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  return string(body), err
}

// --

func (c *Client) GetMembers() (string, error) {
  options := NewClientRequestOptions()
  options.Method = http.MethodGet
  options.Endpoint = "/members"
  return c.MakeRequest(options)
}

func (c *Client) CreateMember(member Member) (string, error) {
  return "", nil
}

func (c *Client) GetMemberById(id string) (string, error) {
  return "", nil
}

func (c *Client) UpdateMemberById(id string) (string, error) {
  return "", nil
}

func (c *Client) DeleteMemberById(id string) (string, error) {
  return "", nil
}
