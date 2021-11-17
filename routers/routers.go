package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
