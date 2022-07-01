package dto

import "github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"

type (
	RegisterEmployeeRequestBody struct {
		Fullname   string `json:"fullname" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
		RoleID     *uint `json:"role_id"`
		DivisionID *uint `json:"division_id" validate:"required"`
	}
	EmployeeResponse struct {
		ID       uint   `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}
	EmployeeWithJWTResponse struct {
		EmployeeResponse
		JWT string `json:"jwt"`
	}
	EmployeeDetailResponse struct {
		EmployeeResponse
		Role     model.Role     `json:"role"`
		Division model.Division `json:"division"`
	}
	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

func (r *RegisterEmployeeRequestBody) FillDefaults() {
	var defaultRoleID uint = 1
	if r.RoleID == nil {
		r.RoleID = &defaultRoleID
	}
}
