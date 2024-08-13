// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	jwt "github.com/dgrijalva/jwt-go"
	mock "github.com/stretchr/testify/mock"
)

// JwtServiceInterface is an autogenerated mock type for the JwtServiceInterface type
type JwtServiceInterface struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: username, isAdmin
func (_m *JwtServiceInterface) GenerateToken(username string, isAdmin bool) (string, error) {
	ret := _m.Called(username, isAdmin)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (string, error)); ok {
		return rf(username, isAdmin)
	}
	if rf, ok := ret.Get(0).(func(string, bool) string); ok {
		r0 = rf(username, isAdmin)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(username, isAdmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateAdmin provides a mock function with given fields: token
func (_m *JwtServiceInterface) ValidateAdmin(token *jwt.Token) bool {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateAdmin")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*jwt.Token) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: token
func (_m *JwtServiceInterface) ValidateToken(token string) (*jwt.Token, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJwtServiceInterface creates a new instance of JwtServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtServiceInterface {
	mock := &JwtServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}