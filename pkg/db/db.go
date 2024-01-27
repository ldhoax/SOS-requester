package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfingDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(cnf ConfingDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
