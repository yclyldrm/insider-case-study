package infrastructure

import (
	"insider-case-study/config"
	"insider-case-study/internal/domain/message"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.GetVar("DB_NAME")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Migrator().DropTable(&message.Message{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&message.Message{}); err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&message.Message{},
	)
}

func FillData(filepath string, db *gorm.DB) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("db didn't find %s", err.Error())
		return
	}
	_, err = sqlDB.Exec(string(file))
	if err != nil {
		log.Printf("Failed to execute SQL script: %v", err)
		return
	}

	log.Println("SQL script executed successfully!")
}
