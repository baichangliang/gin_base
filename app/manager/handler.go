package manager

import (
	"fmt"
	"gin_test/app/models"
	"gin_test/conf"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateManager 管理员创建
func CreateManager(ctx *gin.Context) {
	DB := conf.GetDB()
	var manager models.Manager

	if err := ctx.BindJSON(&manager); err != nil {
		fmt.Println(err.Error())
		Fail(ctx, gin.H{"err": err}, "参数错误")
		return
	}
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(manager.Password), bcrypt.DefaultCost)
	if err != nil {
		Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常, 错误代码(10001)")
		return
	}
	newManager := models.Manager{
		Password: string(hashPassword),
		UserName: manager.UserName,
		Phone:    manager.Phone,
		NickName: manager.NickName,
		Email:    manager.Email,
		Sex:      manager.Sex,
		Avatar:   manager.Avatar,
		Balance:  0,
	}

	tx := DB.Begin()
	if _err := tx.Create(&newManager).Error; _err != nil {
		tx.Rollback()
		Fail(ctx, gin.H{"err": _err}, "参数错误, 请修改后重试.")
		return
	}
	tx.Commit()
	Success(ctx, gin.H{}, "success")
}

// Login 管理员登陆
func Login(ctx *gin.Context) {
	var loginStruct LoginValidators
	err := ctx.BindJSON(&loginStruct)
	if err != nil {
		Fail(ctx, gin.H{"err": err.Error()}, "参数错误")
		return
	}
	PassWord, UserName, DB := loginStruct.Password, loginStruct.UserName, conf.GetDB()
	if PassWord != "" && UserName != "" {
		DB = DB.Where("user_name = ?", UserName)
	} else {
		Fail(ctx, gin.H{}, "缺少必要参数")
		return
	}
	var instance models.Manager
	// 进行查询
	if err := DB.First(&instance).Error; err != nil {
		Fail(ctx, gin.H{}, "用户不存在")
		return
	}
	//判断密码收否正确
	if err := bcrypt.CompareHashAndPassword([]byte(instance.Password), []byte(PassWord)); err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		Fail(ctx, gin.H{}, "密码错误")
		return
	}
	token, err := ReleaseToken(instance)
	if err != nil {
		Fail(ctx, gin.H{}, "系统异常, 错误代码(10002)")
		return
	}
	Success(ctx,
		gin.H{"token": token, "first_name": instance.NickName,
			"phone": instance.Phone, "avatar": instance.Avatar}, "success")
}
