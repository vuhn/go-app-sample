// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccessTokenService is an autogenerated mock type for the AccessTokenService type
type AccessTokenService struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: claims, secretKey
func (_m *AccessTokenService) GenerateToken(claims map[string]interface{}, secretKey string) (string, error) {
	ret := _m.Called(claims, secretKey)

	var r0 string
	if rf, ok := ret.Get(0).(func(map[string]interface{}, string) string); ok {
		r0 = rf(claims, secretKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, string) error); ok {
		r1 = rf(claims, secretKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: token, key
func (_m *AccessTokenService) ValidateToken(token string, key string) (bool, error) {
	ret := _m.Called(token, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(token, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
