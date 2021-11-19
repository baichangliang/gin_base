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
	Name      string `gorm:"type:varchar(30);not null"`
	ManagerID *uint
	Team      []*Genre `gorm:"foreignkey:ManagerID"`
	Goods     []Goods  `gorm:"foreignkey:GenreId"`
}

// Goods 商品
type Goods struct {
	gorm.Model

	Name    string `gorm:"type:varchar(30);not null;unique"`
	GenreId int    `gorm:"not null"`

	Number         int     `gorm:"default:1;comment:'商品编号';unique"`
	Unit           int     `gorm:"default:1;comment:'计量单位:1-瓶,2-盒,3-罐,4-杯,5-包,6-袋'"`
	Standards      string  `gorm:"type:varchar(20);not null;comment:'规格'"`
	MaterialCode   string  `gorm:"type:varchar(50);not null;comment:'物料代码'"`
	MaterialName   string  `gorm:"type:varchar(50);not null;comment:'物料名称'"`
	StandardsPrice float32 `gorm:"default:1;comment:'标准单价'"`
	InsidePrice    float32 `gorm:"default:1;comment:'内部单价'"`
	CompanyPrice   float32 `gorm:"default:1;comment:'公司单价'"`
	Shelf          int     `gorm:"default:1;comment:'是否上架:1-是,0-否'"`
	Rotation       string  `gorm:"type:varchar(512);not null;comment:'轮播图'"`
	Thumbnail      string  `gorm:"type:varchar(200);not null;comment:'缩略图-首页图'"`
	Context        string  `gorm:"type:varchar(500);not null;comment:'商品介绍'"`
}
