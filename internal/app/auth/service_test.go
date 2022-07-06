package auth

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/stretchr/testify/assert"
)

func TestServiceLoginByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx = context.Background()
		payload = dto.ByEmailAndPasswordRequest{
			Email: "vincentlhubbard@superrito.com",
			Password: "123abcABC!",
		}
	)
	res, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
	asserts.NotEmpty(res.JWT)
}

func TestServiceLoginByEmailAndPasswordRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx = context.Background()
		payload = dto.ByEmailAndPasswordRequest{
			Email: "azkaframadhan@superrito.com",
			Password: "123abcABC!",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceLoginByEmailAndPasswordunmatchedEmailAndPassword(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx = context.Background()
		payload = dto.ByEmailAndPasswordRequest{
			Email: "vincentlhubbard@superrito.com",
			Password: "1234567890",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 400")
	}
}

func TestServiceRegisterByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx = context.Background()
		divisionID = uint(1)
		payload = dto.RegisterEmployeeRequestBody{
			Fullname: "Azka Fadhli Ramadhan",
			Email: "azkaframadhan@superrito.com",
			Password: "123abcABC!",
			DivisionID: &divisionID,
		}
	)
	payload.FillDefaults()
	res, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(payload.Fullname, res.Fullname)
	asserts.Equal(payload.Email, res.Email)
	asserts.NotEmpty(res.JWT)
}

func TestServiceRegisterByEmailAndPasswordUserExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx = context.Background()
		divisionID = uint(1)
		payload = dto.RegisterEmployeeRequestBody{
			Fullname: "Vincent L Hubbard",
			Email: "vincentlhubbard@superrito.com",
			Password: "123abcABC!",
			DivisionID: &divisionID,
		}
	)
	_, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}