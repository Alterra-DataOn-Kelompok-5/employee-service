package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/pkg/enum"
	"gorm.io/gorm"
)

func roleSeeder(db *gorm.DB) {
	now := time.Now()
	var roles = []model.Role{
		{
			Name: enum.Role.String(1),
			Common: model.Common{
				ID: 1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Role.String(2),
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
