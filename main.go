package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/urfave/negroni"
)

const (
  NameEnvVar  = "HELLO_NAME"
  DefaultName = "world"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handlerHello)

  n := negroni.Classic()
  n.UseHandler(mux)

  n.Run()
}

func getEnv(envVar string, defaultVal string) string {
  ret := os.Getenv(envVar)

  if ret == "" {
    return defaultVal
  } else {
    return ret
  }
}

func handlerHello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello %s!", getEnv(NameEnvVar, DefaultName))
}
