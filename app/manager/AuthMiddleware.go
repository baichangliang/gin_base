package manager

import (
	"gin_test/app/models"
	"gin_test/conf"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware(auth string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, auth) {
			Fail(ctx, gin.H{}, "未提供认证信息")
			ctx.Abort()
			return
		}
		countSplit := strings.Split(tokenString, " ")[1]
		token, claims, err := ParseToken(countSplit)
		if err != nil || !token.Valid {
			Fail(ctx, gin.H{}, "认证过期, 请重新登陆")
			ctx.Abort()
			return
		}
		ManagerId := claims.ManagerId
		DB := conf.GetDB()
		var m models.Manager
		DB.First(&m, ManagerId)
		// 用户
		if m.ID == 0 {
			Fail(ctx, gin.H{}, "认证错误, 请重新登陆")
			ctx.Abort()
			return
		}
		// 用户存在 将user 的信息写入上下文
		ctx.Set("manager", m)
		ctx.Next()
	}
}
