// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entities

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbAdvert = "tb_advert"

// TbAdvert mapped from table <tb_advert>
type TbAdvert struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Title     *string        `gorm:"column:title;type:varchar(32)" json:"title"`
	Href      *string        `gorm:"column:href;type:longtext" json:"href"`
	Images    *string        `gorm:"column:images;type:longtext" json:"images"`
	IsShow    *int64         `gorm:"column:is_show;type:tinyint(1)" json:"is_show"`
}

// TableName TbAdvert's table name
func (*TbAdvert) TableName() string {
	return TableNameTbAdvert
}
