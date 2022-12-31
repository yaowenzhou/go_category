package repository

import (
	"go_category/domain/model"

	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	FindCategoryById(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryById(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]*model.Category, error)
}

type CategoryRepository struct {
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{}
}
