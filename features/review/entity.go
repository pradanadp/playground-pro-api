package review

import (
	"time"
)

type ReviewCore struct {
<<<<<<< HEAD
	ReviewID  string 
	UserID    string 
	VenueID   string 
	Review    string 
	Rating    float64
	CreatedAt time.Time      
	UpdatedAt time.Time      
	DeletedAt time.Time 
	User      UserCore           
	Venue     VenueCore          
}

type UserCore struct {
	UserID         string                    
	Fullname       string                    
	Email          string                    
	Phone          string                    
	Password       string                    
	Bio            string                    
	Address        string                    
	Role           string                    
	ProfilePicture string                    
	CreatedAt      time.Time                 
	UpdatedAt      time.Time                 
	DeletedAt      time.Time            
	Venues         []VenueCore             
	Reservations   []ReservationCore 
	Reviews        []ReviewCore           
}

type VenueCore struct {
	VenueID       string 
	OwnerID       string 
	Category      string 
	Name          string 
	Description   string 
	Location      string 
	Price         float64
	Longitude     float64
	Latitude      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	VenuePictures []VenuePictureCore
	Reviews       []ReviewCore
}

type VenuePictureCore struct {
	VenuePictureID string
	VenueID        string
	URL            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type ReservationCore struct {
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
=======
	ReviewID  string
	UserID    string
	VenueID   string
	Review    string
	Rating    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	User      UserCore
	Venue     VenueCore
}

type UserCore struct {
	UserID         string
	Fullname       string
	Email          string
	Phone          string
	Password       string
	Bio            string
	Address        string
	Role           string
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Venues         []VenueCore
}

type VenueCore struct {
	VenueID     string
	OwnerID     string
	Category    string
	Name        string
	Description string
	Location    string
	Price       float64
	Longitude   float64
	Latitude    float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Reviews     []ReviewCore
>>>>>>> e733440 (Update entity to core)
}
