package http

import (
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/app/auth"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/app/division"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/app/employee"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/app/role"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	employee.NewHandler(f).Route(v1.Group("/employees"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
	division.NewHandler(f).Route(v1.Group("/divisions"))
	role.NewHandler(f).Route(v1.Group("/roles"))
}
