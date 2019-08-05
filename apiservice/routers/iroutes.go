package routers

import "github.com/labstack/echo/v4"

type ITeoRoutes interface {
	MasterRoutes(server *echo.Echo)
}
