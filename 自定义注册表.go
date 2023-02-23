package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 创建一个自定义注册表
	registry := prometheus.NewRegistry()
	// 可选：添加process和Go运行时指标到我们自定义的注册表中
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector())

	// 创建一个简单的gauge指标
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degree Celsius",
	})

	// 使用我们自定义的注册表注册 gauge
	registry.MustRegister(temp)

	// 设置gague的值为39
	temp.Set(39)

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		Registry: registry,
	}))
	http.ListenAndServe(":8080", nil)
}
