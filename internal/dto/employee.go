package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	EmployeeDetailResponse struct {
		EmployeeResponse
		Role     RoleResponse     `json:"role"`
		Division DivisionResponse `json:"division"`
	}
	// Find employees responses
	// 
	// swagger:model  
	EmployeeResponse struct {
		// required: true
		ID       uint   `json:"id"`
		// required: true
		Fullname string `json:"fullname"`
		// required: true
		Email    string `json:"email"`
	}
	EmployeeWithJWTResponse struct {
		EmployeeResponse
		JWT string `json:"jwt"`
	}
	EmployeeWithCUDResponse struct {
		EmployeeResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
	UpdateEmployeeRequestBody struct {
		ID         *uint   `param:"id" validate:"required"`
		Fullname   *string `json:"fullname" validate:"omitempty"`
		Email      *string `json:"email" validate:"omitempty,email"`
		Password   *string `json:"password" validate:"omitempty"`
		RoleID     *uint   `json:"role_id" validate:"omitempty"`
		DivisionID *uint   `json:"division_id" validate:"omitempty"`
	}
)
