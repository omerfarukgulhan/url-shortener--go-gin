package postgresql

import (
	"gorm.io/gorm"
	"log"
	"url-shortener--go-gin/domain"
)

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Url{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	log.Println("Tables migrated successfully.")
}
