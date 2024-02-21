package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP operations",
		},
		[]string{"method", "endpoint", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime).Seconds()
		endpoint := c.Request.URL.Path
		method := c.Request.Method
		status := c.Writer.Status()
		statusTm := fmt.Sprintf("%d", status)

		//httpRequestsTotal.WithLabelValues(method, endpoint, http.StatusText(status)).Inc()
		//increase(http_requests_total{status="500"}[1m])
		//rate(http_requests_total[1m])
		httpRequestsTotal.WithLabelValues(method, endpoint, statusTm).Inc()
		httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
	}
}

func main() {
	r := gin.Default()

	r.Use(PrometheusMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})
	// 正常的GET请求，返回200状态码
	r.GET("/normal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 故意返回400状态码的GET请求
	r.GET("/bad_request", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	})

	// 故意返回500状态码的GET请求
	r.GET("/server_error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8080")
}
