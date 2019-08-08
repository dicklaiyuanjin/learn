package main

import (
  "net"
  "net/rpc"
  "log"
  "learn/rpc3/hello"
  "net/rpc/jsonrpc"
)

func main() {
  rpc.RegisterName("HelloService", new(hello.HelloService))

  listener, err := net.Listen("tcp", ":1234")
  if err != nil {
    log.Fatal("ListenTCP error: ", err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal("Accept error: ", err)
    }

    go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
  }
}
