package main

import (
  "net"
  "net/rpc"
  "log"
  "learn/protobuf3/hello"
)

func main() {
  server := rpc.NewServer()
  err := hello.RegisterHelloService(server, new(hello.HelloService))
  if err != nil {
    log.Fatal("register error: ", err)
  }

  listener, err := net.Listen("tcp", ":1234")
  if err != nil {
    log.Fatal("ListenTCP error: ", err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal("Accept error: ", err)
    }

    go server.ServeConn(conn)
  }
}
