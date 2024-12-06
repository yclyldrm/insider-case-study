package infrastructure

import (
	"fmt"
	"insider-case-study/config"
	"insider-case-study/internal/domain/message"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.GetVar("DB_HOST"), config.GetVar("DB_PORT"), config.GetVar("DB_USER"), config.GetVar("DB_PASSWORD"), config.GetVar("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&message.Message{},
	)
}
