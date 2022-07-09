package division

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx                 = context.Background()
	divisionService     = NewService(factory.NewFactory())
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestDivisionServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := divisionService.Find(ctx, &testFindAllPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Name)
		asserts.NotEmpty(val.ID)
	}
}
func TestDivisionServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := divisionService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestDivisionServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := divisionService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestDivisionServiceUpdataByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := divisionService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testDivisionName, res.Name)
}

func TestDivisionServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := divisionService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestDivisionServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := divisionService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestDivisionServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	
	asserts := assert.New(t)
	_, err := divisionService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestDivisionServiceCreateDivisionSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	res, err := divisionService.Store(ctx, &testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*testCreatePayload.Name, res.Name)
}

func TestDivisionServiceCreateDivisionAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	_, err := divisionService.Store(ctx, &testCreatePayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
