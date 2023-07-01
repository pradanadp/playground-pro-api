// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	venue "github.com/playground-pro-project/playground-pro-api/features/venue"
	pagination "github.com/playground-pro-project/playground-pro-api/utils/pagination"
	mock "github.com/stretchr/testify/mock"
)

// VenueData is an autogenerated mock type for the VenueData type
type VenueData struct {
	mock.Mock
}

// DeleteVenueImage provides a mock function with given fields: venueImageID
func (_m *VenueData) DeleteVenueImage(venueImageID string) error {
	ret := _m.Called(venueImageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(venueImageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditVenue provides a mock function with given fields: userId, venueId, request
func (_m *VenueData) EditVenue(userId string, venueId string, request venue.VenueCore) error {
	ret := _m.Called(userId, venueId, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, venue.VenueCore) error); ok {
		r0 = rf(userId, venueId, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllVenueImage provides a mock function with given fields: venueID
func (_m *VenueData) GetAllVenueImage(venueID string) ([]venue.VenuePictureCore, error) {
	ret := _m.Called(venueID)

	var r0 []venue.VenuePictureCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]venue.VenuePictureCore, error)); ok {
		return rf(venueID)
	}
	if rf, ok := ret.Get(0).(func(string) []venue.VenuePictureCore); ok {
		r0 = rf(venueID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]venue.VenuePictureCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(venueID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVenueImageByID provides a mock function with given fields: venueID, venueImageID
func (_m *VenueData) GetVenueImageByID(venueID string, venueImageID string) (venue.VenuePictureCore, error) {
	ret := _m.Called(venueID, venueImageID)

	var r0 venue.VenuePictureCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (venue.VenuePictureCore, error)); ok {
		return rf(venueID, venueImageID)
	}
	if rf, ok := ret.Get(0).(func(string, string) venue.VenuePictureCore); ok {
		r0 = rf(venueID, venueImageID)
	} else {
		r0 = ret.Get(0).(venue.VenuePictureCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(venueID, venueImageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertVenue provides a mock function with given fields: userID, venueReq, venueImageReq
func (_m *VenueData) InsertVenue(userID string, venueReq venue.VenueCore, venueImageReq venue.VenuePictureCore) (venue.VenueCore, error) {
	ret := _m.Called(userID, venueReq, venueImageReq)

	var r0 venue.VenueCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, venue.VenueCore, venue.VenuePictureCore) (venue.VenueCore, error)); ok {
		return rf(userID, venueReq, venueImageReq)
	}
	if rf, ok := ret.Get(0).(func(string, venue.VenueCore, venue.VenuePictureCore) venue.VenueCore); ok {
		r0 = rf(userID, venueReq, venueImageReq)
	} else {
		r0 = ret.Get(0).(venue.VenueCore)
	}

	if rf, ok := ret.Get(1).(func(string, venue.VenueCore, venue.VenuePictureCore) error); ok {
		r1 = rf(userID, venueReq, venueImageReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyVenues provides a mock function with given fields: userId
func (_m *VenueData) MyVenues(userId string) ([]venue.VenueCore, error) {
	ret := _m.Called(userId)

	var r0 []venue.VenueCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]venue.VenueCore, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []venue.VenueCore); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]venue.VenueCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterVenue provides a mock function with given fields: userId, request
func (_m *VenueData) RegisterVenue(userId string, request venue.VenueCore) (venue.VenueCore, error) {
	ret := _m.Called(userId, request)

	var r0 venue.VenueCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, venue.VenueCore) (venue.VenueCore, error)); ok {
		return rf(userId, request)
	}
	if rf, ok := ret.Get(0).(func(string, venue.VenueCore) venue.VenueCore); ok {
		r0 = rf(userId, request)
	} else {
		r0 = ret.Get(0).(venue.VenueCore)
	}

	if rf, ok := ret.Get(1).(func(string, venue.VenueCore) error); ok {
		r1 = rf(userId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchVenues provides a mock function with given fields: keyword, latitude, longitude, page
func (_m *VenueData) SearchVenues(keyword string, latitude float64, longitude float64, page pagination.Pagination) ([]venue.VenueCore, int64, int, error) {
	ret := _m.Called(keyword, latitude, longitude, page)

	var r0 []venue.VenueCore
	var r1 int64
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(string, float64, float64, pagination.Pagination) ([]venue.VenueCore, int64, int, error)); ok {
		return rf(keyword, latitude, longitude, page)
	}
	if rf, ok := ret.Get(0).(func(string, float64, float64, pagination.Pagination) []venue.VenueCore); ok {
		r0 = rf(keyword, latitude, longitude, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]venue.VenueCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, float64, float64, pagination.Pagination) int64); ok {
		r1 = rf(keyword, latitude, longitude, page)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, float64, float64, pagination.Pagination) int); ok {
		r2 = rf(keyword, latitude, longitude, page)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(string, float64, float64, pagination.Pagination) error); ok {
		r3 = rf(keyword, latitude, longitude, page)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SelectVenue provides a mock function with given fields: venueId
func (_m *VenueData) SelectVenue(venueId string) (venue.VenueCore, error) {
	ret := _m.Called(venueId)

	var r0 venue.VenueCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (venue.VenueCore, error)); ok {
		return rf(venueId)
	}
	if rf, ok := ret.Get(0).(func(string) venue.VenueCore); ok {
		r0 = rf(venueId)
	} else {
		r0 = ret.Get(0).(venue.VenueCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(venueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterVenue provides a mock function with given fields: userId, venueId
func (_m *VenueData) UnregisterVenue(userId string, venueId string) error {
	ret := _m.Called(userId, venueId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, venueId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VenueAvailability provides a mock function with given fields: venueId
func (_m *VenueData) VenueAvailability(venueId string) (venue.VenueCore, error) {
	ret := _m.Called(venueId)

	var r0 venue.VenueCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (venue.VenueCore, error)); ok {
		return rf(venueId)
	}
	if rf, ok := ret.Get(0).(func(string) venue.VenueCore); ok {
		r0 = rf(venueId)
	} else {
		r0 = ret.Get(0).(venue.VenueCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(venueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVenueData creates a new instance of VenueData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVenueData(t interface {
	mock.TestingT
	Cleanup(func())
}) *VenueData {
	mock := &VenueData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
