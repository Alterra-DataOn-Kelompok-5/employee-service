package division

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/database/seeder"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/mocks"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/pkg/util"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	claims            = util.CreateJWTClaims(testEmail, testEmployeeID, testRoleID, testDivisionID)
	db                = database.GetConnection()
	echoMock          = mocks.EchoMock{E: echo.New()}
	divisionHandler   = NewHandler(&f)
	f                 = factory.Factory{DivisionRepository: repository.NewDivisionRepository(db)}
	testDivisionID    = uint(1)
	testEmail         = "vincentlhubbard@superrito.com"
	testEmployeeID    = uint(1)
	testRoleID        = uint(1)
	testDivisionName  = "Finance Dept."
	testUpdatePayload = dto.UpdateDivisionRequestBody{ID: &testDivisionID, Name: &testDivisionName}
	testCreatePayload = dto.CreateDivisionRequestBody{Name: &testDivisionName}
)

func TestHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/divisions")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}
func TestHandlerGetUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/divisions")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Get(c)) {
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

	c.SetPath("/api/v1/divisions")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestHandlerGetByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	divisionID := "a"

	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.GetById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerGetByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.GetById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerGetByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.GetById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerGetByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.GetById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestHandlerUpdateByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	divisionID := "a"
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerUpdateByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}
func TestHandlerUpdateByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.UpdateById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerUpdateByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.UpdateById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestHandlerDeleteByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	divisionID := "a"
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestHandlerDeleteByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestHandlerDeleteByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))

	c.SetPath("/api/v1/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.DeleteById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerDeleteByIdSuccess(t *testing.T){
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	divisionID := strconv.Itoa(int(testDivisionID))
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/divisions")
	c.SetParamNames("id")
	c.SetParamValues(divisionID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "deleted_at")
	}
}

func TestHandlerCreateInvalidPayload(t *testing.T) {
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(dto.CreateDivisionRequestBody{})
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/divisions")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Create(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestHandlerCreateDivisionAlreadyExist(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	name := "Finance"
	payload, err := json.Marshal(dto.CreateDivisionRequestBody{Name: &name})
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/divisions")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Create(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "duplicate")
	}
}

func TestHandlerCreateUnauthorized(t *testing.T) {
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/divisions")
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Create(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestHandlerCreateSuccess(t *testing.T){
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	token, err := util.CreateJWTToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/divisions")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(divisionHandler.Create(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}
