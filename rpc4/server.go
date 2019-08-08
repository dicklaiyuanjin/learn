package main

import (
  "net/rpc"
  "net/rpc/jsonrpc"
  "net/http"
  "io"

  "learn/rpc4/hello"
)

func main() {
  rpc.RegisterName("HelloService", new(hello.HelloService))

  http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
    var conn io.ReadWriteCloser = struct {
      io.Writer
      io.ReadCloser
    }{
      ReadCloser: r.Body,
      Writer: w,
    }

    rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
  })

  http.ListenAndServe(":1234", nil)
}
