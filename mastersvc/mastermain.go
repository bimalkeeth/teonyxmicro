package main

import (
	"github.com/micro/go-micro/v2"
	"log"
	serviceproto "teonyxmicro/mastersvc/proto/builder"
	api "teonyxmicro/mastersvc/protoapi"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.master_service"),
		micro.Version("latest"),
		micro.AfterStart(serviceStop),
		micro.BeforeStart(serviceStart))

	service.Init()
	err := serviceproto.RegisterMasterServiceHandler(service.Server(), new(api.MasterService))
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func serviceStart() error {

	return nil
}

func serviceStop() error {
	return nil
}
