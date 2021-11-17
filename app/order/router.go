package order

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	// 不需要登陆权限
	ord := e.Group("/api/order/")
	{
		ord.POST("unified/", UnifiedOrder) // 统一下单
	}

}
