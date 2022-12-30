package repository

import "go_category/domain/model"

type CategoryRepository interface {
	InitTable() error
	FindCategoryById(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryById(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]*model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return
}
