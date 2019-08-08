package main

import (
  "net"
  "net/rpc"
  "log"
  "learn/rpc2/hello"
  "learn/rpc2/hsi"
)

func main() {
  hsi.RegisterHelloService(new(hello.HelloService))

  listener, err := net.Listen("tcp", ":1234")
  if err != nil {
    log.Fatal("ListenTCP error: ", err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal("Accept error: ", err)
    }

    go rpc.ServeConn(conn)
  }
}
