package seeder

import (
	"log"
	"time"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/pkg/enum"
	"gorm.io/gorm"
)

func divisionSeeder(db *gorm.DB) {
	now := time.Now()
	var divisions = []model.Division{
		{
			Name: enum.Division.String(1),
			Common: model.Common{
				ID: 1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Division.String(2),
			Common: model.Common{
				ID: 2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Division.String(3),
			Common: model.Common{
				ID: 3,
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
