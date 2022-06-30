package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"gorm.io/gorm"
)

func employeeSeeder(db *gorm.DB) {
	now := time.Now()
	var employees = []model.Employee{
		{
			Fullname: "Vincent L. Hubbard",
			Email: "vincentlhubbard@superrito.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID: 1,
			DivisionID: 1,
			Common: model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "Devon C. Thomas",
			Email: "devoncthomas@superrito.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID: 2,
			DivisionID: 1,
			Common: model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "Bettina M. Easter",
			Email: "bettinameaster@superrito.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID: 2,
			DivisionID: 2,
			Common: model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
		},
	}
	if err := db.Create(&employees).Error; err != nil {
		log.Printf("cannot seed data employees, with error %v\n", err)
	}
	log.Println("success seed data employees")
}
