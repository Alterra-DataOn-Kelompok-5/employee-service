package factory

import (
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
)

type Factory struct {
	EmployeeRepository repository.Employee
	DivisionRepository repository.Division
	RoleRepository     repository.Role
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewEmployeeRepository(db),
		repository.NewDivisionRepository(db),
		repository.NewRoleRepository(db),
	}
}
