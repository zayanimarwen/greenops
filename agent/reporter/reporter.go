package reporter

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type PodMetric struct {
    Name      string
    Namespace string
    CPUUsage  int
    MemUsage  int
}

func SendMetrics(metrics []PodMetric) error {
    jsonData, _ := json.Marshal(metrics)
    resp, err := http.Post("http://backend:8080/metrics", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return nil
}
