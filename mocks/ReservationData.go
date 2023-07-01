// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	time "time"

	reservation "github.com/playground-pro-project/playground-pro-api/features/reservation"
	mock "github.com/stretchr/testify/mock"
)

// ReservationData is an autogenerated mock type for the ReservationData type
type ReservationData struct {
	mock.Mock
}

// CheckAvailability provides a mock function with given fields: venueId
func (_m *ReservationData) CheckAvailability(venueId string) ([]reservation.AvailabilityCore, error) {
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
func (_m *ReservationData) DetailTransaction(userId string, paymentId string) (reservation.PaymentCore, error) {
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

// GetReservationsByTimeSlot provides a mock function with given fields: venueID, checkInDate, checkOutDate
func (_m *ReservationData) GetReservationsByTimeSlot(venueID string, checkInDate time.Time, checkOutDate time.Time) ([]reservation.ReservationCore, error) {
	ret := _m.Called(venueID, checkInDate, checkOutDate)

	var r0 []reservation.ReservationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) ([]reservation.ReservationCore, error)); ok {
		return rf(venueID, checkInDate, checkOutDate)
	}
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) []reservation.ReservationCore); ok {
		r0 = rf(venueID, checkInDate, checkOutDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.ReservationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Time, time.Time) error); ok {
		r1 = rf(venueID, checkInDate, checkOutDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeReservation provides a mock function with given fields: userId, r, p
func (_m *ReservationData) MakeReservation(userId string, r reservation.ReservationCore, p reservation.PaymentCore) (reservation.ReservationCore, reservation.PaymentCore, error) {
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
func (_m *ReservationData) MyReservation(userId string) ([]reservation.MyReservationCore, error) {
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

// MyVenueCharts provides a mock function with given fields: userId, keyword, request
func (_m *ReservationData) MyVenueCharts(userId string, keyword string, request reservation.MyReservationCore) ([]reservation.MyReservationCore, error) {
	ret := _m.Called(userId, keyword, request)

	var r0 []reservation.MyReservationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, reservation.MyReservationCore) ([]reservation.MyReservationCore, error)); ok {
		return rf(userId, keyword, request)
	}
	if rf, ok := ret.Get(0).(func(string, string, reservation.MyReservationCore) []reservation.MyReservationCore); ok {
		r0 = rf(userId, keyword, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservation.MyReservationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, reservation.MyReservationCore) error); ok {
		r1 = rf(userId, keyword, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriceVenue provides a mock function with given fields: venueID
func (_m *ReservationData) PriceVenue(venueID string) (float64, error) {
	ret := _m.Called(venueID)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (float64, error)); ok {
		return rf(venueID)
	}
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(venueID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(venueID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReservationCheckOutDate provides a mock function with given fields: reservation_id
func (_m *ReservationData) ReservationCheckOutDate(reservation_id string) (time.Time, error) {
	ret := _m.Called(reservation_id)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (time.Time, error)); ok {
		return rf(reservation_id)
	}
	if rf, ok := ret.Get(0).(func(string) time.Time); ok {
		r0 = rf(reservation_id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reservation_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReservationStatus provides a mock function with given fields: request
func (_m *ReservationData) ReservationStatus(request reservation.PaymentCore) (reservation.PaymentCore, error) {
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

// NewReservationData creates a new instance of ReservationData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReservationData(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReservationData {
	mock := &ReservationData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
