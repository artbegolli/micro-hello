package main

import (
	"context"
	"fmt"

	"github.com/artbegolli/micro-hello/metadata"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/registry/consul"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *metadata.Request, rsp *metadata.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	fmt.Println("hello world")
	registry := consul.NewRegistry()
	broker := kafka.NewBroker()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()
	service.Run()
}

func runClient(service micro.Service) {
	// Create new greeter client
	greeter := metadata.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &metadata.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
