// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package domain

const TableNameTbArticleTag = "tb_article_tag"

// TbArticleTag mapped from table <tb_article_tag>
type TbArticleTag struct {
	TagID     int64 `gorm:"column:tag_id;type:bigint unsigned;primaryKey" json:"tag_id"`
	ArticleID int64 `gorm:"column:article_id;type:bigint unsigned;primaryKey" json:"article_id"`
}

// TableName TbArticleTag's table name
func (*TbArticleTag) TableName() string {
	return TableNameTbArticleTag
}
