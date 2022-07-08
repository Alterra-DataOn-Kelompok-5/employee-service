package model

type Division struct {
	Name string `json:"name" gorm:"varchar;not_null;unique"`
	Common
}
