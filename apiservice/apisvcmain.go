package main

import (
	"github.com/labstack/echo/v4"
	"teonyxmicro/apiservice/routers"
)
import "github.com/labstack/echo/v4/middleware"

func main() {
	apiServer := echo.New()
	apiServer.Use(middleware.Logger())
	apiServer.Use(middleware.Recover())
	if apiServer == nil {

		apiServer.Logger.Debug("Error in starting server")
	}
	r := routers.New()
	r.MasterRoutes(apiServer)

}
