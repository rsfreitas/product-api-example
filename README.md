# HTTP API example

This repository presents a golang HTTP API example using a protobuf specification
to set both the service API and its database data modeling.

## Example source code organization

The `product-resource` directory stores a golang example application that demonstrates
how to use a protobuf specification to set its API and how to handle it.

For details of this part, consult its proper [documentation](product-resource/README.md).

## Protobuf specification

The example application uses a protobuf file to declare its HTTP API, by defining
its endpoints, request and response messages.

From this protobuf file, a source code is generated that encapsulates the HTTP
framework initialization as well as is endpoints routing. Also, it generates
an OpenAPI documentation of this declared API.

For more details about the protobuf usage here, consult its [documentation](protobuf/README.md).

## Building and testing locally

In order to test this example a few steps must be followed:

1. Put the database server to execute

To ease the setup of the MySQL database server a helper script `database/setup-database.sh`
can be used to install its latest docker version, set up the required credentials,
database and table, and put it to run.

Use the `-i` option to install:
```bash
./database/setup-database.sh -i
```
And the `-s` to put it to run:
```bash
./database/setup-database.sh -s
```

2. Build the application

Inside the directory `product-resource` execute the following command:
```bash
go build
```

3. Execute the application

To execute the application, from within `product-resource` directory, execute:
```bash
./product-resource
```

If everything went well, you'll see similar lines in the output:
```bash
{"time":"2023-12-15T14:07:12.15263-03:00","level":"INFO","msg":"service is running","service.version":"v1.0.0","service.name":"product-resource"}

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```