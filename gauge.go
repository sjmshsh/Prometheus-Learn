package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 创建一个自定义注册表
	registry := prometheus.NewRegistry()

	queueLength := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "queue_length",
		Help: "The number of items in the queue",
	})

	// 使用Set()设置指定的值
	queueLength.Set(0)

	// 增加或者减少
	queueLength.Inc()   // +1: Increment the gauge by 1
	queueLength.Desc()  // -1
	queueLength.Add(23) // Increment by 23
	queueLength.Set(42) // Decrement by 42

	// 使用我们自定义的注册表注册 gauge
	registry.MustRegister(queueLength)

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		Registry: registry,
	}))
	http.ListenAndServe(":8080", nil)
}
