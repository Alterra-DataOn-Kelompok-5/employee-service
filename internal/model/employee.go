package model

type Employee struct {
	Fullname   string `json:"fullname" gorm:"varchar;not_null"`
	Email      string `json:"email" gorm:"varchar;not_null;unique"`
	Password   string `json:"password" gorm:"varchar;not_null"`
	RoleID     uint   `json:"role_id"`
	Role       Role
	DivisionID uint `json:"division_id"`
	Division   Division
	Common
}
