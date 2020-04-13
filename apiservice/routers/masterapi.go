package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro"
	"net/http"
)

var service micro.Service

func init() {
	service = micro.NewService(micro.Name("master_service"))
	service.Init()

}

func New() ITeoRoutes {
	return TeoRoutes{}
}
func (TeoRoutes) MasterRoutes(server *echo.Echo) {
	server.GET("/", getHome)

}
func getHome(context echo.Context) error {

	return context.JSON(http.StatusOK, "Hello World")
}
