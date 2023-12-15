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
