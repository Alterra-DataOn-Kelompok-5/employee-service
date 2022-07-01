package model

import (
	"time"

	"gorm.io/gorm"
)

type Common struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DelatedAt *gorm.DeletedAt `json:"delated_at"`
}
