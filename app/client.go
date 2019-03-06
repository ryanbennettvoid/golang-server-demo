package main

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"

  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo"
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
  Headers  map[string]interface{}
  Body     map[string]interface{}
}

func NewClientRequestOptions() ClientRequestOptions {
  return ClientRequestOptions{
    Headers: map[string]interface{}{
      "Content-Type": "application/json",
    },
    Body: map[string]interface{}{},
  }
}

func (c *Client) MakeRequest(options ClientRequestOptions) ([]byte, error) {
  url := c.BaseUrl + options.Endpoint
  requestBody, err := json.Marshal(options.Body)
  req, err := http.NewRequest(options.Method, url, bytes.NewBuffer(requestBody))
  if err != nil {
    return []byte{}, err
  }
  client := &http.Client{
    Timeout: time.Second * 1,
  }
  res, err := client.Do(req)
  if err != nil {
    return []byte{}, err
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return []byte{}, err
  }
  if res.StatusCode != http.StatusOK {
    errorResponse := ErrorResponse{}
    if err := json.Unmarshal(body, &errorResponse); err != nil {
      fmt.Println(string(body))
      panic(err)
    }
    return []byte{}, errors.New(errorResponse.Message)
  }
  return body, nil
}

// --

func (c *Client) GetHeartbeat() (Heartbeat, error) {
  options := NewClientRequestOptions()
  options.Method = http.MethodGet
  options.Endpoint = "/heartbeat"
  data, err := c.MakeRequest(options)
  if err != nil {
    return Heartbeat{}, err
  }
  var obj Heartbeat
  err = json.Unmarshal(data, &obj)
  if err != nil {
    return Heartbeat{}, err
  }
  return obj, nil
}

func (c *Client) GetMembers() ([]repo.Member, error) {
  options := NewClientRequestOptions()
  options.Method = http.MethodGet
  options.Endpoint = "/members"
  data, err := c.MakeRequest(options)
  if err != nil {
    return []repo.Member{}, err
  }
  var obj []repo.Member
  err = json.Unmarshal(data, &obj)
  if err != nil {
    return []repo.Member{}, err
  }
  return obj, nil
}

func (c *Client) CreateMember(member repo.Member) error {
  options := NewClientRequestOptions()
  options.Method = http.MethodPost
  options.Endpoint = "/members"
  requestMemberData, err := json.Marshal(member)
  if err != nil {
    return err
  }
  err = json.Unmarshal(requestMemberData, &options.Body)
  if err != nil {
    return err
  }
  _, err = c.MakeRequest(options)
  if err != nil {
    return err
  }
  return nil
}

func (c *Client) GetMemberById(id string) (repo.Member, error) {
  options := NewClientRequestOptions()
  options.Method = http.MethodGet
  options.Endpoint = fmt.Sprintf("/members/%s", id)
  data, err := c.MakeRequest(options)
  if err != nil {
    return repo.Member{}, err
  }
  var obj repo.Member
  err = json.Unmarshal(data, &obj)
  if err != nil {
    return repo.Member{}, err
  }
  return obj, nil
}

// func (c *Client) UpdateMemberById(id string) (string, error) {
//   return "", nil
// }

// func (c *Client) DeleteMemberById(id string) (string, error) {
//   return "", nil
// }
