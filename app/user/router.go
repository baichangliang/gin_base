package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/api/", GetWxOpenId)
}
