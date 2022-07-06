package dto

type (
	RegisterEmployeeRequestBody struct {
		Fullname   string `json:"fullname" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
		RoleID     *uint  `json:"role_id"`
		DivisionID *uint  `json:"division_id" validate:"required"`
	}
	UpdateEmployeeRequestBody struct {
		ID         *uint   `param:"id" validate:"required"`
		Fullname   *string `json:"fullname" validate:"omitempty"`
		Email      *string `json:"email" validate:"omitempty,email"`
		Password   *string `json:"password" validate:"omitempty"`
		RoleID     *uint   `json:"role_id" validate:"omitempty"`
		DivisionID *uint   `json:"division_id" validate:"omitempty"`
	}
	EmployeeResponse struct {
		ID       uint   `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}
	RoleResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	DivisionResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	EmployeeWithJWTResponse struct {
		EmployeeResponse
		JWT string `json:"jwt"`
	}
	EmployeeDetailResponse struct {
		EmployeeResponse
		Role     RoleResponse     `json:"role"`
		Division DivisionResponse `json:"division"`
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
