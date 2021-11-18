package goods

import (
	"fmt"
	"gin_test/app/models"
	"gin_test/conf"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// CreateCategory 商品类别
func CreateCategory(ctx *gin.Context) {
	DB := conf.GetDB()
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		fmt.Println(err.Error())
		Fail(ctx, gin.H{"err": err}, "参数错误")
		return
	}

	newCategory := models.Category{
		Name: category.Name,
	}

	tx := DB.Begin()
	if _err := tx.Create(&newCategory).Error; _err != nil {
		tx.Rollback()
		Fail(ctx, gin.H{"err": _err}, "参数错误, 请修改后重试.")
		return
	}
	tx.Commit()
	Success(ctx, gin.H{}, "success")
}

// ListCategory 商品类别列表
func ListCategory(c *gin.Context) {
	categories := make([]models.Category, 0)
	DB := conf.GetDB()
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	// 条件查询--商品类别名称模糊查询
	if Name, isExist := c.GetQuery("Name"); isExist == true && strings.Replace(Name, " ", "", -1) != "" {
		println(Name)
		DB = DB.Where("name LIKE ?", "%"+Name+"%")
	}
	// 分页
	if page > 0 && pageSize > 0 {
		DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	// 总数量
	var total int
	DB.Model(&models.Category{}).Count(&total)
	// 进行查询
	if err := DB.Preload("Friends").Find(&categories).Error; err != nil {
		return
	}
	// 结果返回
	Success(c, gin.H{
		"data": categories, "page": page, "pageSize": pageSize, "count": total,
	}, "success")
}
