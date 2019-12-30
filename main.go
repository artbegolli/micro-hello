package main

import (
	"context"
	"fmt"
	"os"

	"github.com/artbegolli/micro-hello/metadata"
	"github.com/micro/cli"
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

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	if err := metadata.RegisterGreeterHandler(service.Server(), new(Greeter)); err != nil {
		return
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
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
