package employee

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/mocks"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandlerGetInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/employees")
	c.QueryParams().Add("page", "a")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(400, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerGetSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/employees")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
	}
}


func TestHandlerGetByIdInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := "a"
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.GetByID(c)) {
		asserts.Equal(400, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerGetByIdNotFound(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(10)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.GetByID(c)) {
		asserts.Equal(404, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerGetByIdSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(1)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.GetByID(c)) {
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
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := "a"
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerUpdateByIdNotFound(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(10)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerUpdateByIdSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(1)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
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
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := "a"
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerDeleteByIdNotFound(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(10)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerDeleteByIdSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/employees")
	c.SetParamNames("id")
	employeeID := strconv.Itoa(1)
	c.SetParamValues(employeeID)

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	employeeHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(employeeHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)
		
		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "deleted_at")
	}
}
