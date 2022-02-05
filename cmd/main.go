package main

import (
    "log"
    "fmt"
    "net/http"
)

func main() {
    log.Println("go-web-entry")

    http.HandleFunc("/", servHello)
    http.ListenAndServe("0.0.0.0:8080", nil)
}

func servHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
