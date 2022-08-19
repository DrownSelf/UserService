package middlewares

import (
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/pkg/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger(r *repositories.LogRepo) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		log := model.Log{
			LogTime:    params.TimeStamp.Format(time.RFC1123),
			Method:     params.Method,
			Latency:    params.Latency,
			StatusCode: params.StatusCode,
			Url:        params.Request.URL,
		}

		stringLog := fmt.Sprintf("%s|%s|%s|%d \n",
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Latency,
			params.StatusCode,
		)

		if err := r.ReportLog(context.Background(), &log); err != nil {
			fmt.Print(err)
		}
		return stringLog
	})
}
