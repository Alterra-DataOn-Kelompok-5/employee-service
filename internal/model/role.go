package model

type Role struct {
	Name string `json:"name" gorm:"varchar;not_null;unique"`
	Common
}
