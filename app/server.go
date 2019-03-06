package main

const DEFAULT_PORT = 9090

type Server struct {
  Port int
}

func NewServer() Server {
  return Server{
    Port: DEFAULT_PORT,
  }
}
