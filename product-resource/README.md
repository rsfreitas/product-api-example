# product-resource

`product-resource` is the name of a golang example application that demonstrates
how to use a protobuf file to specify its HTTP API (endpoints, inputs and outputs)
and handle it.

It provides a CRUD API to handle "products" (with an ID and a name only), using
a MySQL database to store them.

## Source code organization

All API definition is made through the protobuf file, leaving the source code to
be responsible only for the business logic.

The `main.go` source file is the place where the HTTP is initialized, by setting
the service name and listened port.

The `source.go` source file is where a structure is declared to hold everything
that is required (and important) for the service. This structure also implements
the API declared inside the protobuf file (the `service` block with the RPCs).

Each method corresponds to an RPC declared and is responsible for handling a
request from a client.

The project also has an `internal` directory with a `database` package which
implements the underlying database API for the service.

## API

Its API is declared inside the [project file](../protobuf/proto/services/product_resource/v1/product_resource_api.proto)
and its [OpenAPI documentation](../protobuf/gen/openapi/services/product_resource/v1/openapi.yaml)
can be consulted for details.

> Note: The [swagger editor](https://editor.swagger.io) can be used to render
> the OpenAPI yaml file.

## Building and running the example

In order to build the application, only the go compiler is needed. And it can
be called as:

```bash
go build
```

> Note: The application requires a MySQL database to initialize. You need to make
> sure that it is running and available for the service to execute.

> Note 2: If any change is made at the protobuf declaration, it must be recompiled
> inside the **protobuf** directory. And, an easy way to reflect the changes immediately
> is to edit the `go.mod` file and insert the following line below the `module`
> declaration: `replace github.com/rsfreitas/product-api-example/protobuf => ../buf`.
