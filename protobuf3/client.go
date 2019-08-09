package main

import (
  "log"
  "fmt"

  "learn/protobuf3/hello"
)

func main() {
  client, err := hello.DialHelloService("tcp", "localhost:1234")
  if err != nil {
    log.Fatal("dialing: ", err)
  }

  var reply hello.String
  err = client.Hello(&hello.String{Value: "hello"}, &reply)
  if err != nil {
    log.Fatal("calling: ", err)
  }

  fmt.Println(reply)
}
