package database

import (
	"erbeyinn/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func Connect() error {
	var err error

	dsn := "host=localhost user=postgres password=gans356pb786 dbname=dbtest sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		return err
	}
	_, err = db.DB()

	if err != nil {
		return err
	}

	return nil

}

func AutoMigrater() error{
	err := db.AutoMigrate(&models.Book{})
	if err!=nil{
		return err
	}

	return nil
}