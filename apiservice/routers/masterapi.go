package routers

import (
	"github.com/labstack/echo/v4"
)

func New() ITeoRoutes {
	return TeoRoutes{}
}

func (TeoRoutes) MasterRoutes(server *echo.Echo) {

}
