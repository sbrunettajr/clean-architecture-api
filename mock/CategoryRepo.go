// Code generated by mockery v2.10.0. DO NOT EDIT.

package mock

import (
	entity "github.com/sbrunettajr/clean-architecture-api/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// CategoryRepo is an autogenerated mock type for the CategoryRepo type
type CategoryRepo struct {
	mock.Mock
}

// AddCategory provides a mock function with given fields: category
func (_m *CategoryRepo) AddCategory(category entity.Category) error {
	ret := _m.Called(category)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Category) error); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCategories provides a mock function with given fields:
func (_m *CategoryRepo) GetCategories() ([]entity.Category, error) {
	ret := _m.Called()

	var r0 []entity.Category
	if rf, ok := ret.Get(0).(func() []entity.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryByUUID provides a mock function with given fields: categoryUUID
func (_m *CategoryRepo) GetCategoryByUUID(categoryUUID string) (entity.Category, error) {
	ret := _m.Called(categoryUUID)

	var r0 entity.Category
	if rf, ok := ret.Get(0).(func(string) entity.Category); ok {
		r0 = rf(categoryUUID)
	} else {
		r0 = ret.Get(0).(entity.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
