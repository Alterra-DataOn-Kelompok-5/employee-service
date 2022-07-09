package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateRoleRequestBody struct {
		Name *string `json:"name" validate:"required"`
	}
	UpdateRoleRequestBody struct {
		ID   *uint   `param:"id" validate:"required"`
		Name *string `json:"name" validate:"required"`
	}
	RoleResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	RoleWithCUDResponse struct {
		RoleResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
