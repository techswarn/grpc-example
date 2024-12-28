package main

import (
//  "fmt"
  "io"
  "net/http"
 // "os"
  "strings"

  "golang.org/x/net/http2"
  "golang.org/x/net/http2/h2c"
  "google.golang.org/grpc"
)


func main() {
  g := grpc.NewServer()
  // gRPC server setup...

  handler := func(w http.ResponseWriter, r *http.Request) {
    if r.ProtoMajor == 2 {
      if strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
        g.ServeHTTP(w, r)
        return
      }
      io.WriteString(w, "Hello HTTP/2")
      return
    }
    io.WriteString(w, "Hello HTTP/1.1")
  }

  server := &http.Server{
//    Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
	Addr: ":8080",
    Handler: h2c.NewHandler(http.HandlerFunc(handler), &http2.Server{}),
  }
  server.ListenAndServe()
}