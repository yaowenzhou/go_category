package repository

import (
	"fmt"
	"go_category/domain/model"

	"gorm.io/gorm"
)

// CategoryRepositoryInterface TODO
type CategoryRepositoryInterface interface {
	FindCategoryById(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryById(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategorysByLevel(int32) ([]*model.Category, error)
	FindCategorysByParent(int64) ([]*model.Category, error)
}

// CategoryRepository TODO
type CategoryRepository struct {
	db *gorm.DB
}

// FindCategoryById 通过id查询Category
func (s *CategoryRepository) FindCategoryById(id int64) (*model.Category, error) {
	return model.GetCategoryById(s.db, id)
}

// CreateCategory 创建Category
func (s *CategoryRepository) CreateCategory(info *model.Category) (int64, error) {
	return model.CreateCategory(info, s.db)
}

// DeleteCategoryById 删除Category
func (s *CategoryRepository) DeleteCategoryById(id int64) error {
	return model.DeleteCategoryById(id, s.db)
}

// UpdateCategory 更新Category
func (s *CategoryRepository) UpdateCategory(info *model.Category) error {
	return model.UpdateCategory(info, s.db)
}

// FindAll 获取所有的Category
func (s *CategoryRepository) FindAll() ([]*model.Category, error) {
	return model.FindAll(&model.Category{}, s.db)
}

// FindCategoryByName TODO
func (s *CategoryRepository) FindCategoryByName(name string) (*model.Category, error) {
	categorys, err := model.FindAll(&model.Category{Name: name}, s.db)
	if err != nil {
		return nil, err
	}
	if len(categorys) == 0 {
		return nil, fmt.Errorf("unable to find category by name(%s)", name)
	}
	return categorys[0], nil
}

// FindCategoryByLevel TODO
func (s *CategoryRepository) FindCategorysByLevel(level int32) ([]*model.Category, error) {
	return model.FindAll(&model.Category{Level: level}, s.db)
}

// FindCategoryByParent TODO
func (s *CategoryRepository) FindCategorysByParent(parent int64) ([]*model.Category, error) {
	return model.FindAll(&model.Category{Parent: parent}, s.db)
}

// NewCategoryRepository TODO
func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{db: db}
}
