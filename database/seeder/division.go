package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"gorm.io/gorm"
)

func divisionSeeder(db *gorm.DB) {
	now := time.Now()
	var divisions = []model.Division{
		{
			Name: "Finance",
			Common: model.Common{
				ID: 1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: "IT",
			Common: model.Common{
				ID: 2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&divisions).Error; err != nil {
		log.Printf("cannot seed data divisions, with error %v\n", err)
	}
	log.Println("success seed data divisions")
}
