package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/server"
	_ "github.com/micro/go-plugins/broker/nats/v2"
	_ "github.com/micro/go-plugins/registry/nats/v2"
	_ "github.com/micro/go-plugins/transport/nats/v2"
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
