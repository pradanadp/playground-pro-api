package reservation

import "time"

type ReservationEntity struct {
	ReservationID string
	UserID        string
	VenueVenueID  string
	PaymentID     string
	CheckInDate   time.Time
	CheckOutDate  time.Time
	Duration      uint
	Subtotal      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
