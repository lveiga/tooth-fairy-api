package patient

import (
	"time"
)

//Patient - model that represents a patient
type Patient struct {
	ID        uint32    `gorm:"primary_key;auto_increment;" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name" binding:"required"`
	Phone     string    `gorm:"size:255;not null;" json:"phone" binding:"required"`
	Email     string    `gorm:"size:100;not null;unique;" json:"email" binding:"required"`
	Age       int       `gorm:"size:3;not null;" json:"age" binding:"required"`
	Gender    string    `gorm:"size:255;not null;" json:"gender" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
