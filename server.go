package main

import (
  "fmt"
  "net/http"
  "github.com/golang/glog"
)

type HasHandleFunc interface { //this is just so it would work for gorilla and http.ServerMux
    HandleFunc(pattern string, handler func(w http.ResponseWriter, req *http.Request))
}
type Handler struct {
    http.HandlerFunc
    Enabled bool
}

type Handlers map[string]*Handler

func (h Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    if handler, ok := h[path]; ok && handler.Enabled {
        handler.ServeHTTP(w, r)
    } else {
        http.Error(w, "Swagger UI Not Found", http.StatusNotFound)
    }
}

func (h Handlers) HandleFunc(mux HasHandleFunc, pattern string, handler http.HandlerFunc) {
    h[pattern] = &Handler{handler, true}
    mux.HandleFunc(pattern, h.ServeHTTP)
}

func hello_world() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World")
  }
}

func startServer() {
  http.HandleFunc("/hello_world", helloWorld())
  http.ListenAndServe(":8080", nil)
  glog.Flush()
}
