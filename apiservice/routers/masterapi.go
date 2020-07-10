package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2"
	master "teonyxmicro/apiservice/services/master"
)

var service micro.Service

func init() {
	service = micro.NewService(micro.Name("master_service"))
	service.Init()

}
func New() IRoutes {
	return Routes{}
}
func (Routes) MasterRoutes(server *echo.Echo) {
	routes := master.New(service)
	server.GET("/", routes.GetHome)
}
