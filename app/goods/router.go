package goods

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	adminGoods := e.Group("/v1/admin/")
	{
		adminGoods.POST("genre/", CreateGenre)       // 商品类别创建
		adminGoods.GET("genre/", ListGenre)          // 商品类别列类
		adminGoods.GET("genre/:ID/", DetailsGenre)   // 商品类别详情
		adminGoods.PUT("genre/:ID/", UpdateGenre)    // 商品类别更新
		adminGoods.DELETE("genre/:ID/", DeleteGenre) // 商品类别删除
	}

}
