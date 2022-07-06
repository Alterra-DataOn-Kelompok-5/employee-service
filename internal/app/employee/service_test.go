package employee

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/stretchr/testify/assert"
)

func TestServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		payload         = dto.SearchGetRequest{}
	)
	res, err := employeeService.Find(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Email)
		asserts.NotEmpty(val.Fullname)
		asserts.NotEmpty(val.ID)
	}
}

func TestServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		payload         = dto.ByIDRequest{ID: 1}
	)
	res, err := employeeService.FindByID(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(uint(1), res.ID)
}

func TestServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	var (
		asserts         = assert.New(t)
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		payload         = dto.ByIDRequest{ID: 4}
	)
	_, err := employeeService.FindByID(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceUpdataByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts         = assert.New(t)
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		id              = uint(1)
		fullname        = "Vincent Luis Hubbard"
		email           = "vincentluishubbars@superrito.com"
		divisionID      = uint(2)
		roleID          = uint(2)
		payload         = dto.UpdateEmployeeRequestBody{
			ID:         &id,
			Fullname:   &fullname,
			Email:      &email,
			DivisionID: &divisionID,
			RoleID:     &roleID,
		}
	)
	res, err := employeeService.UpdateById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(fullname, res.Fullname)
	asserts.Equal(email, res.Email)
	asserts.Equal("IT", res.Division.Name)
	asserts.Equal("User", res.Role.Name)
}

func TestServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	var (
		asserts         = assert.New(t)
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		id              = uint(1)
		fullname        = "Vincent Luis Hubbard"
		email           = "vincentluishubbars@superrito.com"
		divisionID      = uint(2)
		roleID          = uint(2)
		payload         = dto.UpdateEmployeeRequestBody{
			ID:         &id,
			Fullname:   &fullname,
			Email:      &email,
			DivisionID: &divisionID,
			RoleID:     &roleID,
		}
	)
	_, err := employeeService.UpdateById(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts         = assert.New(t)
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		id              = uint(1)
		payload         = dto.ByIDRequest{ID: id}
	)
	res, err := employeeService.DeleteById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts         = assert.New(t)
		employeeService = NewService(factory.NewFactory())
		ctx             = context.Background()
		id              = uint(10)
		payload         = dto.ByIDRequest{ID: id}
	)
	_, err := employeeService.DeleteById(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}
