package main

import (
  "net/rpc"
  "log"
  "fmt"

  "learn/protobuf2/hello"
)

func main() {
  client, err := rpc.Dial("tcp", "localhost:1234")
  if err != nil {
    log.Fatal("dialing: ", err)
  }

  var reply hello.String
  err = client.Call("HelloService.Hello", &hello.String{Value: "hello"}, &reply)
  if err != nil {
    log.Fatal("calling: ", err)
  }

  fmt.Println(reply)
}
