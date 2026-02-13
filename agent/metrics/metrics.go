package metrics

import (
    "k8s.io/client-go/kubernetes"
    v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodMetric struct {
    Name      string
    Namespace string
    CPUUsage  int
    MemUsage  int
}

func CollectPodMetrics(clientset *kubernetes.Clientset) []PodMetric {
    pods, _ := clientset.CoreV1().Pods("").List(v1.ListOptions{})
    metrics := []PodMetric{}
    for _, pod := range pods.Items {
        metrics = append(metrics, PodMetric{
            Name:      pod.Name,
            Namespace: pod.Namespace,
            CPUUsage:  0, // placeholder
            MemUsage:  0, // placeholder
        })
    }
    return metrics
}
