package response

import (
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
)

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}