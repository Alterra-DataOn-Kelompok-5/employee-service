package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/mocks"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandlerLoginByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)
	
	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Invalid parameters or payload","info": null},"error": "bad_request"}`, body)
	}
}

func TestHandlerLoginByEmailAndPasswordUnmatchedEmailAndPassword(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "vincentlhubbard@superrito.com",
		Password: "1234567890",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Email or password is incorrect","info": null},"error": "bad_request"}`, body)
	}
}

func TestHandlerLoginByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "vincentlhubbard@superrito.com",
		Password: "123abcABC!",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}

func TestHandlerRegisterByEmailAndPasswordUserAlreadyExist(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Vincent L. Hubbard"
		email = "vincentlhubbard@superrito.com"
		password = "123abcABC!"
		divisionID = uint(1)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname: fullname,
		Email:    email,
		Password: password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Created value already exists")
	}
}

func TestHandlerRegisterByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Vincent L. Hubbard"
		password = "123abcABC!"
		divisionID = uint(3)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname: fullname,
		Password: password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestHandlerRegisterByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Azka"
		email = "azka@superrito.com"
		password = "123abcABC!"
		divisionID = uint(2)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname: fullname,
		Email: email,
		Password: password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}
