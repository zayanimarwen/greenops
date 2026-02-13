package main

import (
    "fmt"
    "greenops/agent/kube_client"
    "greenops/agent/metrics"
    "greenops/agent/reporter"
)

func main() {
    fmt.Println("GreenOps Agent started")

    // Connect to cluster
    clientset, err := kube_client.GetClient()
    if err != nil {
        panic(err)
    }

    // Collect metrics
    podMetrics := metrics.CollectPodMetrics(clientset)

    // Send to backend
    err = reporter.SendMetrics(podMetrics)
    if err != nil {
        fmt.Println("Error sending metrics:", err)
    }

    fmt.Println("Metrics sent successfully")
}
