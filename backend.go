package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Server"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "Hello API"}`))
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api", apiHandler)
    
    fmt.Println("Server berjalan di http://localhost:8090")
    http.ListenAndServe(":8090", nil)
}
