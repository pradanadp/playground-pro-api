// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	user "github.com/playground-pro-project/playground-pro-api/features/user"
	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: userID
func (_m *UserData) DeleteByID(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateOTP provides a mock function with given fields: request
func (_m *UserData) GenerateOTP(request user.UserCore) (user.UserCore, error) {
	ret := _m.Called(request)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (user.UserCore, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: userID
func (_m *UserData) GetByID(userID string) (user.UserCore, error) {
	ret := _m.Called(userID)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.UserCore, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) user.UserCore); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: request
func (_m *UserData) Login(request user.UserCore) (user.UserCore, string, error) {
	ret := _m.Called(request)

	var r0 user.UserCore
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (user.UserCore, string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) string); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(user.UserCore) error); ok {
		r2 = rf(request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: request
func (_m *UserData) Register(request user.UserCore) (user.UserCore, error) {
	ret := _m.Called(request)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (user.UserCore, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: userID, updatedUser
func (_m *UserData) UpdateByID(userID string, updatedUser user.UserCore) error {
	ret := _m.Called(userID, updatedUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, user.UserCore) error); ok {
		r0 = rf(userID, updatedUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserData(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
