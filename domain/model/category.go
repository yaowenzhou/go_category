package model

import (
	"errors"
	"go_category/pkg"

	"gorm.io/gorm"
)

const CategoryTableName = "category"

type Category struct {
	ID          int64  `gorm:"column:id;primaryKey;not null;autoIncrement:false" json:"id"`
	Name        string `gorm:"column:name;uniqueIndex:,where:is_deleted=false;size:64;not null" json:"name"`
	Level       int32  `gorm:"column:level;not null" json:"level"`
	Parent      int64  `gorm:"column:parent;not null" json:"parent"`
	Image       string `gorm:"column:image;not null" json:"image"`
	Description string `gorm:"column:description;not null; default:''" json:"description"`
	IsDelete    *bool  `gorm:"column:is_delete;not null" json:"is_deleted"`
}

// TableName 获取表名
func (u *Category) TableName() string {
	return CategoryTableName
}

// FindAll 获取所有的数据
func FindAll(info *Category, db *gorm.DB) ([]*Category, error) {
	var datas []*Category
	db = db.Model(info).Where("is_delete = ?", true).Find(&datas)
	return datas, db.Error
}

// GetCategoryById 通过id查询数据
func GetCategoryById(db *gorm.DB, id int64) (*Category, error) {
	var data *Category
	db = db.Model(&Category{ID: id}).Where("is_delete = ?", true).First(&data)
	return data, db.Error
}

// FindCategoryByName 通过name查询数据
func FindCategoryByName(db *gorm.DB, name string) ([]*Category, error) {
	var datas []*Category
	db = db.Model(&Category{Name: name}).Where("is_delete = ?", true).Find(&datas)
	return datas, db.Error
}

// FindCategoryByParent 通过parent查询数据
func FindCategoryByParent(db *gorm.DB, parent int64) ([]*Category, error) {
	var datas []*Category
	db = db.Model(&Category{Parent: parent}).Where("is_delete = ?", true).Find(&datas)
	return datas, db.Error
}

// FindCategorysByLevel 通过level查询数据
func FindCategorysByLevel(db *gorm.DB, level int32) ([]*Category, error) {
	var datas []*Category
	db = db.Model(&Category{Level: level}).Where("is_delete = ?", true).Find(&datas)
	return datas, db.Error
}

// BeforeCreate TODO
func (s *Category) BeforeCreate(db *gorm.DB) (err error) {
	if s.IsDelete == nil {
		s.IsDelete = new(bool)
	}
	return
}

// CreateCategory TODO
func CreateCategory(info *Category, db *gorm.DB) (id int64, err error) {
	if info.ID == 0 {
		info.ID = pkg.GetSnowFlake()
	}
	db = db.Model(info).Create(info)
	return info.ID, db.Error
}

// UpdateCategory TODO
func UpdateCategory(info *Category, db *gorm.DB) (err error) {
	if info.ID == 0 {
		return errors.New("update operation must specify an id")
	}
	db = db.Model(info).Updates(info)
	return db.Error
}

// DeleteCategoryById Tombstone
func DeleteCategoryById(id int64, db *gorm.DB) (err error) {
	isDeleted := true
	db = db.Model(&Category{ID: id}).Update("id", &isDeleted)
	return db.Error
}
