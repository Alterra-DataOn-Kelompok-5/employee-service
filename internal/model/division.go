package model

type Division struct {
	DivisionName string `json:"division_name" gorm:"varchar;not_null;unique"`
	Common
}
