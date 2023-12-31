// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	time "time"

	reservation "github.com/playground-pro-project/playground-pro-api/features/reservation"
	mock "github.com/stretchr/testify/mock"
)

// ReservationService is an autogenerated mock type for the ReservationService type
type ReservationService struct {
	mock.Mock
}

// CheckAvailability provides a mock function with given fields: venueId
func (_m *ReservationService) CheckAvailability(venueId string) ([]reservation.AvailabilityCore, error) {
	ret := _m.Called(venueId)

	var r0 []reservation.AvailabilityCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]reservation.AvailabilityCore, error)); ok {
		return rf(venueId)
	}
	if rf, ok := ret.Get(0).(func(string) []reservation.AvailabilityCore); ok {
		r0 = rf(venueId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.AvailabilityCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(venueId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetailTransaction provides a mock function with given fields: userId, paymentId
func (_m *ReservationService) DetailTransaction(userId string, paymentId string) (reservation.PaymentCore, error) {
	ret := _m.Called(userId, paymentId)

	var r0 reservation.PaymentCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (reservation.PaymentCore, error)); ok {
		return rf(userId, paymentId)
	}
	if rf, ok := ret.Get(0).(func(string, string) reservation.PaymentCore); ok {
		r0 = rf(userId, paymentId)
	} else {
		r0 = ret.Get(0).(reservation.PaymentCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, paymentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeReservation provides a mock function with given fields: userId, r, p
func (_m *ReservationService) MakeReservation(userId string, r reservation.ReservationCore, p reservation.PaymentCore) (reservation.ReservationCore, reservation.PaymentCore, error) {
	ret := _m.Called(userId, r, p)

	var r0 reservation.ReservationCore
	var r1 reservation.PaymentCore
	var r2 error
	if rf, ok := ret.Get(0).(func(string, reservation.ReservationCore, reservation.PaymentCore) (reservation.ReservationCore, reservation.PaymentCore, error)); ok {
		return rf(userId, r, p)
	}
	if rf, ok := ret.Get(0).(func(string, reservation.ReservationCore, reservation.PaymentCore) reservation.ReservationCore); ok {
		r0 = rf(userId, r, p)
	} else {
		r0 = ret.Get(0).(reservation.ReservationCore)
	}

	if rf, ok := ret.Get(1).(func(string, reservation.ReservationCore, reservation.PaymentCore) reservation.PaymentCore); ok {
		r1 = rf(userId, r, p)
	} else {
		r1 = ret.Get(1).(reservation.PaymentCore)
	}

	if rf, ok := ret.Get(2).(func(string, reservation.ReservationCore, reservation.PaymentCore) error); ok {
		r2 = rf(userId, r, p)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MyReservation provides a mock function with given fields: userId
func (_m *ReservationService) MyReservation(userId string) ([]reservation.MyReservationCore, error) {
	ret := _m.Called(userId)

	var r0 []reservation.MyReservationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]reservation.MyReservationCore, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []reservation.MyReservationCore); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.MyReservationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyVenueCharts provides a mock function with given fields: userId, keyword, checkInDate, checkOutDate
func (_m *ReservationService) MyVenueCharts(userId string, keyword string, checkInDate time.Time, checkOutDate time.Time) ([]reservation.MyReservationCore, error) {
	ret := _m.Called(userId, keyword, checkInDate, checkOutDate)

	var r0 []reservation.MyReservationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, time.Time, time.Time) ([]reservation.MyReservationCore, error)); ok {
		return rf(userId, keyword, checkInDate, checkOutDate)
	}
	if rf, ok := ret.Get(0).(func(string, string, time.Time, time.Time) []reservation.MyReservationCore); ok {
		r0 = rf(userId, keyword, checkInDate, checkOutDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.MyReservationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, time.Time, time.Time) error); ok {
		r1 = rf(userId, keyword, checkInDate, checkOutDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReservationStatus provides a mock function with given fields: request
func (_m *ReservationService) ReservationStatus(request reservation.PaymentCore) (reservation.PaymentCore, error) {
	ret := _m.Called(request)

	var r0 reservation.PaymentCore
	var r1 error
	if rf, ok := ret.Get(0).(func(reservation.PaymentCore) (reservation.PaymentCore, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(reservation.PaymentCore) reservation.PaymentCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(reservation.PaymentCore)
	}

	if rf, ok := ret.Get(1).(func(reservation.PaymentCore) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReservationService creates a new instance of ReservationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReservationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReservationService {
	mock := &ReservationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
