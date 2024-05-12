package service

import (
	"errors"
	"learning-golang/go-unit-testing/mock/entity"
	"learning-golang/go-unit-testing/mock/repository"
)

type CategoryService struct {
	category repository.CategoryRepository
}

func (service *CategoryService) GetCategoryById(id int) (*entity.Category, error) {
	category, err := service.category.FindById(id)
	if err != nil {
		return nil, errors.New("CATEGORY_NOT_FOUND")
	}
	return category, nil
}
