package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 创建带house和room标签的gauge指标对象
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// 指定标签名称
		[]string{"house", "root"},
	)

	// 注册到全局默认注册表中
	prometheus.MustRegister(temp)

	// 针对不通标签值设置不通的指标值
	temp.WithLabelValues("cnych", "living-room").Set(27)
	temp.WithLabelValues("cnych", "bedroom").Set(25.3)
	temp.WithLabelValues("ydzs", "living-root").Set(24.5)
	temp.WithLabelValues("ydzs", "bedroom").Set(27.7)

	// 暴露自定义的指标
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
