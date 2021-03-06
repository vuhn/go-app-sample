// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Password is an autogenerated mock type for the Password type
type Password struct {
	mock.Mock
}

// CompareHashAndPassword provides a mock function with given fields: plainPassword, hashedPassword
func (_m *Password) CompareHashAndPassword(plainPassword string, hashedPassword string) bool {
	ret := _m.Called(plainPassword, hashedPassword)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(plainPassword, hashedPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GenerateFromPassword provides a mock function with given fields: plainPassword
func (_m *Password) GenerateFromPassword(plainPassword string) (string, error) {
	ret := _m.Called(plainPassword)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plainPassword)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(plainPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
