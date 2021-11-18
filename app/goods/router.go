package goods

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	ord := e.Group("/api/goods/")
	{
		ord.POST("genre/", CreateGenre) // 商品类别创建
		ord.GET("genre/", ListGenre)    // 商品类别列类
	}

}
