package goods

import (
	"gin_test/app/models"
	"gin_test/conf"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// CreateGenre 商品类别创建
func CreateGenre(ctx *gin.Context) {
	DB := conf.GetDB()
	var genre models.Genre

	if err := ctx.BindJSON(&genre); err != nil {
		Fail(ctx, gin.H{"err": err}, "参数错误")
		return
	}
	// 事务管理
	tx := DB.Begin()
	if _err := tx.Create(&genre).Error; _err != nil {
		tx.Rollback()
		Fail(ctx, gin.H{"err": _err}, "参数错误, 请修改后重试.")
		return
	}
	tx.Commit()
	Success(ctx, gin.H{}, "success")
}

// ListGenre 商品类别列表
func ListGenre(c *gin.Context) {
	genres := make([]models.Genre, 0)
	DB := conf.GetDB()
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	// 条件查询--商品类别名称模糊查询
	if Name, isExist := c.GetQuery("Name"); isExist == true && strings.Replace(Name, " ", "", -1) != "" {
		DB = DB.Where("name LIKE ?", "%"+Name+"%")
	} else {
		DB = DB.Where("manager_id is null || manager_id = '' || manager_id = 0")
	}

	// 分页
	if page > 0 && pageSize > 0 {
		DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	// 总数量
	var total int
	DB.Model(&models.Genre{}).Count(&total)
	// 进行查询.Preload("Friends")
	if err := DB.Preload("Team").Find(&genres).Error; err != nil {
		return
	}
	// 结果返回
	Success(c, gin.H{
		"data": genres, "page": page, "pageSize": pageSize, "count": total,
	}, "success")
}
