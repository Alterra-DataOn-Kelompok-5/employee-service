package model

import "time"

type Common struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DelatedAt time.Time `json:"delated_at"`
}
