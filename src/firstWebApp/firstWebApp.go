package main

import (
  "net/http"
  "fmt"
)

func HelloResponse(rw http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(rw, "First goland App.")
}

func main() {
 http.HandleFunc("/",HelloResponse)
 http.ListenAndServe(":3000",nil)

}
