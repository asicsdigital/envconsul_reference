package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/urfave/negroni"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello %s!", getEnv("HELLO_NAME", "world"))
  })

  n := negroni.Classic()
  n.UseHandler(mux)

  http.ListenAndServe(":3000", n)
}

func getEnv(envVar string, defaultVal string) string {
  ret := os.Getenv(envVar)

  if ret == "" {
    return defaultVal
  } else {
    return ret
  }
}
