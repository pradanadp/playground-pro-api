package data

import (
	"time"

	"github.com/playground-pro-project/playground-pro-api/features/reservation"
	paymentgateway "github.com/playground-pro-project/playground-pro-api/utils/payment_gateway"
	"gorm.io/gorm"
)

type Reservation struct {
	ReservationID string `gorm:"primaryKey;type:varchar(45)"`
	UserID        string `gorm:"foreignKey:UserID;type:varchar(45)"`
	VenueID       string `gorm:"foreignKey:VenueID;type:varchar(45)"`
	PaymentID     *string
	CheckInDate   time.Time `gorm:"type:datetime"`
	CheckOutDate  time.Time `gorm:"type:datetime"`
	Duration      uint
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Payment struct {
	PaymentID     string `gorm:"primaryKey;type:varchar(45)"`
	VANumber      string `gorm:"type:varchar(225);not null"`
	PaymentMethod string
	PaymentType   string
	GrandTotal    string
	ServiceFee    float64
	Status        string         `gorm:"type:enum('pending','success','cancelled','expired');default:'pending'"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Reservation   Reservation    `gorm:"foreignKey:PaymentID;references:PaymentID"`
}

// Reservation-Model to reservation-core
func reservationModels(r Reservation) reservation.ReservationCore {
	return reservation.ReservationCore{
		ReservationID: r.ReservationID,
		UserID:        r.UserID,
		VenueID:       r.VenueID,
		CheckInDate:   r.CheckInDate,
		CheckOutDate:  r.CheckOutDate,
		Duration:      r.Duration,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		DeletedAt:     r.DeletedAt.Time,
	}
}

// Reservation-core to Reservation-Model
func reservationEntities(r reservation.ReservationCore) Reservation {
	return Reservation{
		ReservationID: r.ReservationID,
		UserID:        r.UserID,
		VenueID:       r.VenueID,
		CheckInDate:   r.CheckInDate,
		CheckOutDate:  r.CheckOutDate,
		Duration:      r.Duration,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		DeletedAt:     gorm.DeletedAt{Time: r.DeletedAt},
	}
}

// Payment-Model to payment-core
func paymentModels(p Payment) reservation.PaymentCore {
	return reservation.PaymentCore{
		PaymentID:     p.PaymentID,
		VANumber:      p.VANumber,
		PaymentMethod: p.PaymentMethod,
		PaymentType:   p.PaymentType,
		GrandTotal:    p.GrandTotal,
		ServiceFee:    p.ServiceFee,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// Payment-core to payment-model
func paymentEntities(p reservation.PaymentCore) *Payment {
	return &Payment{
		PaymentID:     p.PaymentID,
		VANumber:      p.VANumber,
		PaymentMethod: p.PaymentMethod,
		PaymentType:   p.PaymentType,
		GrandTotal:    p.GrandTotal,
		ServiceFee:    p.ServiceFee,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// Payment response Midtrans to payment-core
func paymentCoreFromChargeResponse(res *paymentgateway.ChargeResponse) reservation.PaymentCore {
	var vaNumber string
	if len(res.VaNumbers) > 0 {
		vaNumber = res.VaNumbers[0].VANumber
	}

	return reservation.PaymentCore{
		PaymentID:     res.TransactionID,
		VANumber:      vaNumber,
		PaymentMethod: res.PaymentType,
		PaymentType:   res.Bank,
		GrandTotal:    res.GrossAmount,
		ServiceFee:    0,
		Status:        res.TransactionStatus,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		Reservation:   reservation.ReservationCore{},
	}
}

// `gorm:"type:enum('none','card','bca','bri','bni','mandiri','qris','gopay','shopeepay');default:'none'"`
// `gorm:"type:enum('cash','debit_card','bank_transfer','e-wallet');default:'cash'"`
