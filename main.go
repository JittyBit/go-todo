package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", greet)
  http.ListenAndServe(":8080", mux)
}
