package main

import (
    "fmt"
    "log"
    "net/http"
    "greenops/backend/api"
)

func main() {
    fmt.Println("GreenOps Backend started")
    http.HandleFunc("/metrics", api.MetricsHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
