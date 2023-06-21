package database

import (
	"fmt"

	venue_picture "github.com/playground-pro-project/playground-pro-api/features/image/data"
	payment "github.com/playground-pro-project/playground-pro-api/features/payment/data"
	reservation "github.com/playground-pro-project/playground-pro-api/features/reservation/data"
	review "github.com/playground-pro-project/playground-pro-api/features/review/data"
	user "github.com/playground-pro-project/playground-pro-api/features/user/data"
	venue "github.com/playground-pro-project/playground-pro-api/features/venue/data"

	"github.com/playground-pro-project/playground-pro-api/app/config"
	"github.com/playground-pro-project/playground-pro-api/app/middlewares"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var log = middlewares.Log()

func InitDatabase(c *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DBUSER, c.DBPASSWORD, c.DBHOST, c.DBPORT, c.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.SetupJoinTable(&payment.Payment{}, "Cart", &reservation.Reservation{})
	if err != nil {
		fmt.Println("setup error ", err)
		return nil
	}

	err = db.AutoMigrate(
		&user.User{},
		&venue.Venue{},
		&venue_picture.VenuePicture{},
		&payment.Payment{},
		&review.Review{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("success connected and migrated to database")

	initSuperAdmin(db)

	return db
}
