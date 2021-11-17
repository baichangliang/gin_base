package models

import "github.com/jinzhu/gorm"

// 但不存在的时候查询就不会返回结构体中的这个字段`json:"createTime,omitempty" gorm:"createTime,omitempty"`
// json:"-"

// Manager 管理员
type Manager struct {
	gorm.Model
	NickName string  `gorm:"size:128;comment:'用户真实名称'"`
	UserName string  `gorm:"size:128;comment:'用户帐号'"`
	Password string  `gorm:"size:128;comment:'密码'"`
	Avatar   string  `gorm:"size:256;comment:'头像地址'"`
	Email    string  `gorm:"size:128;"`
	Sex      int     `gorm:"default:1;comment:'性别(1-男,2-女,3-保密,0-未知)'"`
	Phone    string  `gorm:"comment:'手机号'"`
	Balance  float32 `gorm:"comment:'余额'"`
}
