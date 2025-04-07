package database

import (
	"time"

	"github.com/google/uuid"
)

type Shorturl struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Original_url string    `gorm:"not null" json:"original_url"`
	Short_url    string    `gorm:"not null" json:"short_url"`
	Created_at   time.Time `gorm:"created_at" json:"created_at"`
}
