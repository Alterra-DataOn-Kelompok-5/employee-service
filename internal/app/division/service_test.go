package division

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx             = context.Background()
	divisionService = NewService(factory.NewFactory())
)

func TestServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.SearchGetRequest{}
	)

	res, err := divisionService.Find(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 2)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Name)
		asserts.NotEmpty(val.ID)
	}
}
func TestServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	res, err := divisionService.FindByID(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	_, err := divisionService.FindByID(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceUpdataByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		name    = "Finance Dept."
		payload = dto.UpdateDivisionRequestBody{
			ID:   &id,
			Name: &name,
		}
	)
	res, err := divisionService.UpdateById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(name, res.Name)
}

func TestServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		name    = "Finance Dept."
		payload = dto.UpdateDivisionRequestBody{
			ID:   &id,
			Name: &name,
		}
	)

	_, err := divisionService.UpdateById(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	res, err := divisionService.DeleteById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(10)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	_, err := divisionService.DeleteById(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestServiceCreateDivisionSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		name = "Finance Dept."
		payload = dto.CreateDivisionRequestBody{
			Name: &name,
		}
	)

	res, err := divisionService.Store(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*payload.Name, res.Name)
}

func TestServiceCreateDivisionAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		name = "Finance"
		payload = dto.CreateDivisionRequestBody{
			Name: &name,
		}
	)

	_, err := divisionService.Store(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
