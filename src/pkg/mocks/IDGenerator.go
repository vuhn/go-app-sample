// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IDGenerator is an autogenerated mock type for the IDGenerator type
type IDGenerator struct {
	mock.Mock
}

// GenerateNewID provides a mock function with given fields:
func (_m *IDGenerator) GenerateNewID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}