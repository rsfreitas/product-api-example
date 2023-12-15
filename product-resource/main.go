package main

import (
	"context"

	"github.com/rsfreitas/go-pocket-utils/logger"

	product_resourcev1 "github.com/rsfreitas/product-api-example/protobuf/gen/go/services/product_resource/v1"
)

func main() {
	serviceName := "product-resource"
	svc := NewService(serviceName)
	httpServer, err := product_resourcev1.NewHttpServer(&product_resourcev1.HttpServerOptions{
		Port:        8080,
		ServerKind:  product_resourcev1.EchoServer,
		Logger:      svc.logger,
		ServiceName: serviceName,
		ApiHandlers: svc,
	})
	if err != nil {
		svc.logger.Fatal(context.TODO(), "could not initialize Service", logger.Error(err))
	}

	httpServer.Run()
}
