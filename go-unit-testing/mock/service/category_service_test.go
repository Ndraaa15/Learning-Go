package service

import (
	"learning-golang/go-unit-testing/mock/entity"
	"learning-golang/go-unit-testing/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{category: categoryRepository}

func TestCategoryService(t *testing.T) {
	c := &entity.Category{ID: 1, Name: "Electronic"}

	categoryRepository.Mock.On("FindById", 1).Return(c, nil)

	category, err := categoryService.GetCategoryById(1)
	assert.NotNil(t, category)
	assert.Nil(t, err)
	assert.Equal(t, c.ID, category.ID)
	assert.Equal(t, c.Name, category.Name)
}
