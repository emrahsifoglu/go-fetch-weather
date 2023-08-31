package main

import (
    "fmt"
    "net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    name := query.Get("name")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}

func main() {
    http.HandleFunc("/hello", handleRequest)
    http.ListenAndServe(":9090", nil)
}
