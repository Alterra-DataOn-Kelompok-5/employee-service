package employee

import (
	"context"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/pkg/enum"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	"github.com/stretchr/testify/assert"
)

var (
	testID                    = uint(1)
	testFullname              = "Vincent Luis Hubbard"
	ctx                       = context.Background()
	testEmployeeService       = NewService(factory.NewFactory())
	testUpdateEmployeePayload = dto.UpdateEmployeeRequestBody{
		ID:         &testID,
		Fullname:   &testFullname,
		Email:      &testEmail,
		DivisionID: &testDivisionID,
		RoleID:     &testAdminRoleID,
	}
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestEmployeeServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testEmployeeService.Find(ctx, &testFindAllPayload)
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

func TestEmployeeServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testEmployeeService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(uint(1), res.ID)
}

func TestEmployeeServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testEmployeeService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestEmployeeServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testEmployeeService.UpdateById(ctx, &testUpdateEmployeePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testFullname, res.Fullname)
	asserts.Equal(testEmail, res.Email)
	asserts.Equal(enum.Division(testDivisionID).String(), res.Division.Name)
	asserts.Equal(enum.Role(testAdminRoleID).String(), res.Role.Name)
}

func TestEmployeeServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testEmployeeService.UpdateById(ctx, &testUpdateEmployeePayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestEmployeeServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testEmployeeService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestEmployeeServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testEmployeeService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}
