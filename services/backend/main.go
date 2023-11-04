package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = 0

func main() {

	prometheus.MustRegister(pingCounter)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		pingCounter.Inc()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/metrics", prometheusHandler())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by ping handler",
	},
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
