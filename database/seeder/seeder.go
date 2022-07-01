package seeder

import (
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) All() {
	s.DB.Exec("DELETE FROM employees")
	s.DB.Exec("DELETE FROM divisions")
	s.DB.Exec("DELETE FROM roles")
	roleSeeder(s.DB)
	divisionSeeder(s.DB)
	employeeSeeder(s.DB)
}