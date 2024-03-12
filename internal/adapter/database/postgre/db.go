package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(dsn string) *DB {
	connection := connect(dsn)
	return &DB{
		Connection: connection,
	}
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&CalendarModel{},
	)
	if err != nil {
		panic(err)
	}
	return db
}
