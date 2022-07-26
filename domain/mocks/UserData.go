// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "backend/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// DeleteUserData provides a mock function with given fields: userid
func (_m *UserData) DeleteUserData(userid int) bool {
	ret := _m.Called(userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RegisterData provides a mock function with given fields: newuser
func (_m *UserData) RegisterData(newuser domain.User) domain.User {
	ret := _m.Called(newuser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newuser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// SearchUserData provides a mock function with given fields: username
func (_m *UserData) SearchUserData(username string) (domain.User, error) {
	ret := _m.Called(username)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserData provides a mock function with given fields: newuser
func (_m *UserData) UpdateUserData(newuser domain.User) domain.User {
	ret := _m.Called(newuser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newuser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}