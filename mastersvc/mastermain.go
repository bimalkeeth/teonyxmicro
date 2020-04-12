package main

import (
	"github.com/micro/go-micro"
	"log"
)
import serviceproto "teonyxmicro/mastersvc/proto/builder"
import api "teonyxmicro/mastersvc/protoapi"

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.master_service"),
		micro.Version("latest"),
		micro.AfterStart(serviceStop),
		micro.BeforeStart(serviceStart))

	err := serviceproto.RegisterMasterServiceHandler(service.Server(), new(api.MasterService))
	if err != nil {
		log.Fatal(err)
	}
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func serviceStart() error {

}

func serviceStop() error {

}
