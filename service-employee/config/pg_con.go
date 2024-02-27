package config

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"service-employee/model"

	"time"
)

func InitPostgres() *gorm.DB {
	var dsn string = "host=localhost user=postgres password=postgres dbname=ms-db port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error while connecting to database")
	}

	db.AutoMigrate(&model.Employee{})
	//additional:
	//set min pool-max pool
	return db
}

func NewPostgreContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)
}
