package repository

import "learning-golang/go-unit-testing/mock/entity"

type CategoryRepository interface {
	FindById(id int) (*entity.Category, error)
}
