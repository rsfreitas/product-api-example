#!/bin/bash

install_database() {
    echo "Install mysql using docker"
    docker pull mysql/mysql-server:latest
}

run_database() {
    echo "Put the database in execution"
    docker run --env="MYSQL_ROOT_PASSWORD=password" --env="MYSQL_USER=user" --env="MYSQL_PASSWORD=password" --env="MYSQL_DATABASE=example" --publish 3306:3306 --name=mysql -d mysql/mysql-server:latest

    echo "Waiting for the server to be up"
    for ((i = 10; i >= 1; i--)); do
        echo "$i seconds remaining..."
        sleep 1
    done

    echo "Copy SQL script for database adjustments"
    docker cp ./database/database.sql mysql:/
    echo "Run the SQL script"
    docker exec mysql /bin/sh -c 'mysql -uroot -ppassword example < /database.sql'
}

usage() {
    echo "Usage: setup-database.sh [OPTIONS]"
    echo "Set up the database for the product-resource service."
    echo
    echo "Options:"
    echo -e " -h\t\tShows this help screen."
    echo -e " -i\t\tInstall MySQL docker image."
    echo -e " -s\t\tPut the database to run."
    echo
}

while getopts :hsi opt; do
    case $opt in
        h)
            usage
            exit 1
            ;;

        i)
            install_database
            ;;

        s)
            run_database
            ;;

        ?)
            echo $opt
            echo "Unsupported option"
            exit 2
            ;;
    esac
done

exit 0
