package middleware

import (
	"strconv"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/metric"
	"github.com/gin-gonic/gin"
)

func Metrics(service metric.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		appMetric := metric.NewHTTP(c.Request.URL.Path, c.Request.Method)
		appMetric.Started()
		c.Next()
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(c.Writer.Status())
		service.SaveHTTP(appMetric)
	}
}
