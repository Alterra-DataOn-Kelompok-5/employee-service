package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewares(e *echo.Echo) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	path := dirname + "/logs/"

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	fileName := path + util.Getenv("LOG_FILE", "employee-service.logs")
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_rfc3339}] ${status} ${method} ${host}${uri} ${latency_human}` + "\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
		Output:           f,
	}))
}

func JWTMiddleware(claims dto.JWTClaims, signingKey []byte) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &dto.JWTClaims{},
		SigningKey: signingKey,
	}
	return middleware.JWTWithConfig(config)
}
