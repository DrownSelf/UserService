package repositories

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "http",
		Name:        "server_rps",
		Help:        "Histogram of response latency of http handlers",
		ConstLabels: nil,
		Buckets:     prometheus.DefBuckets,
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.MustRegister(HttpHistogram)
}

type handlerPath struct {
	sync.Map
}

func (hp *handlerPath) Get(handler string) string {
	v, ok := hp.Load(handler)
	if !ok {
		return ""
	}
	return v.(string)
}

func (hp *handlerPath) set(r gin.RouteInfo) {
	hp.Store(r.Handler, r.Path)
}

type PrometheusRepository struct {
	Engine  *gin.Engine
	Path    *handlerPath
	Updated bool
}

func NewMetricsRepo(e *gin.Engine) *PrometheusRepository {
	if e == nil {
		return nil
	}
	r := PrometheusRepository{
		Engine: e,
		Path:   &handlerPath{},
	}

	return &r
}

func (r *PrometheusRepository) UpdatePath() {
	r.Updated = true
	for _, route := range r.Engine.Routes() {
		r.Path.set(route)
	}
}
