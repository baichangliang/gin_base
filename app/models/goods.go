package models

import "github.com/jinzhu/gorm"

// Category 多对多自引用
type Category struct {
	gorm.Model
	Name    string      `gorm:"size:20;comment:'类别名称'"`
	Friends []*Category `json:"Parent,omitempty" gorm:"many2many:friendships;association_jointable_foreignkey:friend_id,omitempty"`
}

// Genre 一对多自引用(商品类别)
type Genre struct {
	gorm.Model
	Name      string
	ManagerID *uint
	Team      []*Genre `gorm:"foreignkey:ManagerID"`
}
