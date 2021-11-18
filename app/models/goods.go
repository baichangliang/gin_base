package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name    string      `gorm:"size:20;comment:'类别名称'"`
	Friends []*Category `json:"Parent,omitempty" gorm:"many2many:friendships;association_jointable_foreignkey:friend_id,omitempty"`
}
