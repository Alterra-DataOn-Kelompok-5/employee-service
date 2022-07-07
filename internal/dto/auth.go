package dto

import "github.com/golang-jwt/jwt/v4"

type (
	RegisterEmployeeRequestBody struct {
		Fullname   string `json:"fullname" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
		RoleID     *uint  `json:"role_id"`
		DivisionID *uint  `json:"division_id" validate:"required"`
	}

	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		UserID     uint   `json:"user_id"`
		Email      string `json:"email"`
		RoleID     uint   `json:"role_id"`
		DivisionID uint   `json:"division_id"`
		jwt.RegisteredClaims
	}
)

func (r *RegisterEmployeeRequestBody) FillDefaults() {
	var defaultRoleID uint = 1
	if r.RoleID == nil {
		r.RoleID = &defaultRoleID
	}
}
