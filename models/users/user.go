package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID           string  `gorm:"primaryKey" json:"id"`
	FirstName    string  `json:"name"`
	LastName     string  `json:"last_name"`
	Username     *string `gorm:"unique;default:null" json:"username"` // nullable & unique
	Email        string  `json:"email" gorm:"unique"`
	PasswordHash string  `json:"-"` // Exclude from JSON responses
	Role         string  `json:"role" gorm:"default:administration"`
	CreatedAt    int64   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt    int64   `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var lastUser User

	// Find last created user by sorting ID in descending order
	if err := tx.Order("id DESC").First(&lastUser).Error; err != nil {
		// If no user found (first user), set ID to TV0001
		if err == gorm.ErrRecordNotFound {
			u.ID = "TV0001"
			return nil
		}
		return err // Return other DB errors
	}

	// Extract numeric part from last ID (e.g., "TV0007" -> 7)
	var lastNumber int
	fmt.Sscanf(lastUser.ID, "TV%04d", &lastNumber)

	// Increment and format as TV000X with leading zeros
	newNumber := lastNumber + 1
	u.ID = fmt.Sprintf("TV%04d", newNumber)

	return nil
}
