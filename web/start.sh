#!/usr/bin/env bash

docker build -t franko-web .
docker run -it --rm -p 8080:8080 --name franko-web franko-web