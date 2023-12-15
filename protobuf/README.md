# protobuf

This directory holds the example application HTTP API declaration using a protobuf
file to specify it.

In order to encapsulate some of the complexity used by an HTTP server application,
it uses some `protobuf` plugins to do some "dirt" jobs.

* [protoc-gen-httpserver](https://github.com/rsfreitas/protoc-gen-httpserver): this
plugin is responsible for translating the protobuf declarations in some important
parts for the application. It generates all server boilerplate initialization, setting
endpoint routes and handlers for the HTTP server part. It is also responsible for
generating 2 data modelling layers for the application. One database layer (structures
with `Model` suffix) and another for HTTP responses (structures with `HttpResponse`
suffix).

* [protoc-gen-pocket-extensions](https://github.com/rsfreitas/protoc-gen-pocket-extensions):
this plugin is responsible for generating the OpenAPI documentation of the HTTP
API declared.

## Generating the sources

A helper scripts, `generate.sh`, is located in this directory to help installing
required tools for generating the source code from the protobuf files.

In order to use the plugins and compile the protobuf file, the [buf](https://buf.build)
is required. And to install it, this script can be used in the following form:

```bash
./generate.sh -i
```

With `buf` installed the protobuf can be compiled and the sources generated using
the following command:

```bash
./generate.sh
```
