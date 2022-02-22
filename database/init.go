package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	connectionString :=
		"host=localhost " +
			"user=televend " +
			"password=televend " +
			"dbname=televend " +
			"port=5432 " +
			"sslmode=disable " +
			"TimeZone=Europe/Zagreb"
	db, _ := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return db
}
