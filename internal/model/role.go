package model

type Role struct {
	RoleName string `json:"role_name" gorm:"varchar;not_null;unique"`
	Common
}
