package main 
 
import ( 
 "net/http" 
 
 "github.com/prometheus/client_golang/prometheus/promhttp" 
) 
 
func main() { 
    // Serve the default Prometheus metrics registry over HTTP on /metrics. 
 http.Handle("/metrics", promhttp.Handler()) 
 http.ListenAndServe(":8080", nil) 
} 
