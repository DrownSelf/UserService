package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"InnoTaxi/internal/app/repositories"
)

func ObserveStats(r *repositories.PrometheusRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !r.Updated {
			r.UpdatePath()
		}

		start := time.Now()
		c.Next()

		repositories.HttpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			r.Path.Get(c.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}
