package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type PodMetric struct {
    Name      string
    Namespace string
    CPUUsage  int
    MemUsage  int
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var metrics []PodMetric
    err := json.NewDecoder(r.Body).Decode(&metrics)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Printf("Received %d pod metrics\n", len(metrics))
    // TODO: calcul ROI / CO₂
    w.WriteHeader(http.StatusOK)
}
