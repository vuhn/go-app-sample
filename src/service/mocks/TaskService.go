// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/vuhn/go-app-sample/entity"
)

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: _a0
func (_m *TaskService) CreateTask(_a0 *entity.Task) (*entity.Task, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Task
	if rf, ok := ret.Get(0).(func(*entity.Task) *entity.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasksList provides a mock function with given fields: limit, offset
func (_m *TaskService) GetTasksList(limit int, offset int) ([]*entity.Task, int64, error) {
	ret := _m.Called(limit, offset)

	var r0 []*entity.Task
	if rf, ok := ret.Get(0).(func(int, int) []*entity.Task); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Task)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int) int64); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
