package factory

import (
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
)

type Factory struct {
	EmployeeRepository repository.Employee
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewEmployeeRepository(db),
	}
}
