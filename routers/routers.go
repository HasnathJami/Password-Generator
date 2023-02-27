package routers

import (
	"net/http"

	"github.com/HasnathJami/go-echo-redis/controllers"
	"github.com/labstack/echo"
)

func Router() *echo.Echo {
  router := echo.New()

  router.GET("/", func(c echo.Context) error {return c.String(http.StatusOK, "Welcome to Password Generator")})
  router.POST("/generateStrongPassword/:password", controllers.GenerateStrongPassword)
  router.GET("/getStrongPassword", controllers.GetStrongPassword)

  return router
}