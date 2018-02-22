#!/usr/bin/env bash

docker build -t franko-auth .
docker run -it --rm -p 13500:13500 --name franko-auth franko-auth