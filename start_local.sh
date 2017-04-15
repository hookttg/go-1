#!/usr/bin/env bash

# pull docker containers
#docker pull zhanat87/golang
#docker pull postgres
# stop local postgres
sudo service postgresql stop
# link postgres to golang and start containers
docker run --name postgres -d -p 5432:5432 -e POSTGRES_PASSWORD="postgres" -e POSTGRES_DB="go_restful" postgres
docker run -d -p 8080:8080 --name restful --link postgres:postgres zhanat87/golang
#docker run -it --rm -p 8080:8080 --name restful --link postgres:postgres zhanat87/golang
#docker run -it -p 8080:8080 --name restful --link postgres:postgres zhanat87/golang bash