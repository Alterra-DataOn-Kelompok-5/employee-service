package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateDivisionRequestBody struct {
		Name *string `json:"name" validate:"required"`
	}
	UpdateDivisionRequestBody struct {
		ID   *uint   `param:"id" validate:"required"`
		Name *string `json:"name" validate:"required"`
	}
	DivisionResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	DivisionWithCUDResponse struct {
		DivisionResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
