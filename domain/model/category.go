package model

import "gorm.io/gorm"

const CategoryTableName = "category"

type Category struct {
	ID       int64 `gorm:"column:id;primaryKey;not null;autoIncrement:false" json:"id"`
	IsDelete *bool `gorm:"column:is_delete;not null" json:"is_deleted"`
}

// TableName 获取表名
func (u *Category) TableName() string {
	return CategoryTableName
}

// GetCategorys 获取所有的数据
func GetCategorys(db *gorm.DB, info *Category) ([]*Category, error) {
	var datas []*Category
	db = db.Model(&Category{}).Where(info).Where("is_delete = ?", true).Find(&datas)
	return datas, db.Error
}

// BeforeCreate TODO
func (s *Category) BeforeCreate(db *gorm.DB) (err error) {
	if s.IsDelete == nil {
		s.IsDelete = new(bool)
	}
	return
}
