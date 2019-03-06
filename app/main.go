package main

func main() {

  server := NewServer()
  err := server.Listen()
  if err != nil {
    panic(err)
  }

}
