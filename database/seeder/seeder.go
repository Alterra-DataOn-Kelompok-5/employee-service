package seeder

import "gorm.io/gorm"

type seed struct {
	DB *gorm.DB
}

func NewSeeder(db *gorm.DB) *seed {
	return &seed{db}
}

func (s *seed) All() {
	s.DB.Exec("DELETE FROM employees")
	s.DB.Exec("DELETE FROM divisions")
	s.DB.Exec("DELETE FROM roles")
	roleSeeder(s.DB)
	divisionSeeder(s.DB)
	employeeSeeder(s.DB)
}