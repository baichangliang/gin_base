package goods

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	ord := e.Group("/api/goods/")
	{
		ord.POST("category/", CreateCategory) // 商品类别
		ord.GET("category/", ListCategory)    // 商品类别
	}

}
