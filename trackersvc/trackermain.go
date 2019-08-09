package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
	"log"
	"net"
	"time"
)

func opts(o *micro.Options) {
	o.Server = server.NewServer(func(o *server.Options) {
		o.Name = "teonyx.tracker"
		o.Broker = nats.NewBroker()
	})
}

func main() {
	err := cmd.Init()
	if err != nil {
		log.Fatal("error in command initialization")
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}
	for _, add := range addrs {
		log.Println(add.Network()+":", add.String())
	}

	microServer := micro.NewService(opts)
	if serr := microServer.Run(); serr != nil {

		log.Fatal("Error in server running")
	}
	retry := time.NewTicker(1 * time.Second)

RetryLoop:
	for {
		select {
		case <-retry.C:
			if err = microServer.Options().Broker.Connect(); err != nil {
				log.Fatal("error in broker connection")
			} else {
				retry.Stop()
				break RetryLoop
			}
		}
	}
	if err = microServer.Run(); err != nil {
		log.Fatal("Error in server running")
	}
}
