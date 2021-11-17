package manager

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	// 不需要登陆权限
	unauthorized := e.Group("/api/manager/")
	{
		unauthorized.POST("", CreateManager) // 管理员创建
		unauthorized.POST("/login/", Login)  // 管理员登陆
	}

	// pc端, 需要公司人员认证权限
	pcUser := e.Group("/api/user_pc/", AuthMiddleware("Bearer "))
	{
		pcUser.GET("", CreateManager)
	}
}
