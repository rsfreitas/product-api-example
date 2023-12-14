#!/bin/bash

PATH=$PATH:$HOME/go/bin
PATH=$PATH:/usr/local/go/bin

# Here we download/install everything we need to compile/generate code of all
# protobuf files of this repository.
plugins_install() {
    if ! test -d plugins; then
        mkdir plugins
    fi

    if ! test -d plugins/protoc-gen-pocket-extensions; then
        (cd plugins && git clone https://github.com/rsfreitas/protoc-gen-pocket-extensions.git)
        (cd plugins/protoc-gen-pocket-extensions && go generate ./options/pocket && go build && go install)
    fi

    if ! test -d plugins/protoc-gen-httpserver; then
        (cd plugins && git clone https://github.com/rsfreitas/protoc-gen-httpserver.git)
        (cd plugins/protoc-gen-httpserver && go generate ./options/pocket && go build && go install)
    fi
}

usage() {
    echo "Usage: generate.sh [OPTIONS]"
    echo "Generate go source from protobuf files."
    echo
    echo "Options:"
    echo -e " -h\t\tShows this help screen."
    echo -e " -i\t\tInstall dependencies like buf compiler. You may need sudo."
    echo
    echo "The default behavior is to execute this script without arguments."
    echo "This way all protobuf files will be compiled."
    echo
    echo "Examples:"
    echo
    echo "$ ./generate.sh"
    echo
    echo "> It will generate go source for all protobuf files inside"
    echo "'gen/go/services' folder."
    echo
}

install_deps() {
    BIN="/usr/local/bin" && \
    VERSION="1.17.0" && \
    curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
        -o "${BIN}/buf" && \

    chmod +x "${BIN}/buf"
    buf --version
}

while getopts :hs:uilf opt; do
    case $opt in
        h)
            usage
            exit 1
            ;;

        i)
            install_deps
            exit 1
            ;;

        ?)
            echo $opt
            echo "Unsupported option"
            exit -1
            ;;
    esac
done

plugins_install
rm -rf gen
buf generate proto
cd gen/go/services || (echo "could not find generated files"; exit 1)

echo "New files generated ✅🚀"
exit 0

