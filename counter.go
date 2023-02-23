func main() {
	// 创建一个自定义注册表
	registry := prometheus.NewRegistry()

	totalRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of handled HTTP requests.",
	})

	totalRequests.Inc()   // +1
	totalRequests.Add(23) // +n

	// 使用我们自定义的注册表注册 gauge
	registry.MustRegister(totalRequests)

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		Registry: registry,
	}))
	http.ListenAndServe(":8080", nil)
}
