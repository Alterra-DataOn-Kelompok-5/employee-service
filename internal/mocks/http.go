package mocks

import (
	"io"
	"net/http/httptest"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = &util.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
