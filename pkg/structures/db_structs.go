package structures

import (
	"time"
)

type DbConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOSTNAME string
	DB_NAME     string
}

type ProjectInfo struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Category    string    `gorm:"not null" json:"category"`
	Description string    `gorm:"not null" json:"description"`
	StartDate   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"start_date"`
	Status      string    `gorm:"not null" json:"status"`
	EndDate     time.Time `gorm:"default:null" json:"end_date"`
}
