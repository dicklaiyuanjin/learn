package main

import (
  "log"
  "fmt"

  "learn/rpc2/hsi"
)

func main() {
  client, err := DialHelloService("tcp", "localhost:1234")
  if err != nil {
    log.Fatal("dialing: ", err)
  }

  var reply string
  err = client.Hello("hello", &reply)
  if err != nil {
    log.Fatal("calling: ", err)
  }

  fmt.Println(reply)
}
