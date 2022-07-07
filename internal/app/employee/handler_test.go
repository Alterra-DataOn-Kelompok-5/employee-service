package employee

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/mocks"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/pkg/util"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	claims          = util.CreateJWTClaims(testEmail, testEmployeeID, testRoleID, testDivisionID)
	db              = database.GetConnection()
	echoMock        = mocks.EchoMock{E: echo.New()}
	employeeHandler = NewHandler(&f)
	f               = factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	testDivisionID  = uint(1)
	testEmail       = "vincentlhubbard@superrito.com"
	testRoleID      = uint(1)
	testEmployeeID  = uint(1)
)

func TestHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerGetUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/employees")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerGetSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c.SetPath("/api/v1/employees")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
	}
}

func TestHandlerGetByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	employeeID := "a"

	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.GetById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerGetByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.GetById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerGetByUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.GetById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerGetByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.GetById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "role")
		asserts.Contains(body, "division")
	}
}

func TestHandlerUpdateByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	employeeID := "a"
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerUpdateByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerUpdateByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerUpdateByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "role")
		asserts.Contains(body, "division")
	}
}

func TestHandlerDeleteByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	employeeID := "a"
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerDeleteByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerDeleteByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))

	c.SetPath("/api/v1/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerDeleteByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	employeeID := strconv.Itoa(int(testEmployeeID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/employees")
	c.SetParamNames("id")
	c.SetParamValues(employeeID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "deleted_at")
	}
}
