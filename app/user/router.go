package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	// 不需要登陆权限
	unauthorized := e.Group("/api/user_un/")
	{
		unauthorized.GET("", GetWxOpenId) //
	}

	// pc端, 需要公司人员认证权限
	pcUser := e.Group("/api/user_pc/", AuthMiddleware("Bearer "))
	{
		pcUser.GET("", GetWxOpenId)
	}
}
