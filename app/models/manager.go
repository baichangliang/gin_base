package models

import "github.com/jinzhu/gorm"

type Manager struct { // 核销人员
	// 登陆帐号 真实姓名 密码 是否禁用 手机 邮箱 角色 头像 配送地址
	gorm.Model
	FirstName string `gorm:"size:128;"`
	Head      string `gorm:"size:256;comment:'头像地址'"`
	Email     string `gorm:"size:128;"`
	Sex       int    `gorm:"default:5;comment:'性别(1-男,2-女,3-保密,0-未知)'"`
	Phone     string `gorm:"size:11;comment:'手机号'"`
}
