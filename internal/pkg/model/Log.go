package model

import (
	"net/url"
	"time"
)

type Log struct {
	LogTime    string        `bson:"timeStamp"`
	Method     string        `bson:"method"`
	StatusCode int           `bson:"statusCode"`
	Latency    time.Duration `bson:"latency"`
	Url        *url.URL      `bson:"url"`
}
