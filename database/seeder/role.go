package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"gorm.io/gorm"
)

func roleSeeder(db *gorm.DB) {
	now := time.Now()
	var roles = []model.Role{
		{
			RoleName: "Admin",
			Common: model.Common{
				ID: 1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			RoleName: "User",
			Common: model.Common{
				ID: 2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&roles).Error; err != nil {
		log.Printf("cannot seed data roles, with error %v\n", err)
	}
	log.Println("success seed data roles")
}
