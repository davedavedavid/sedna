#!/bin/bash

docker build -f mistnet-yolo-server.Dockerfile -t decshub.org/mistnet-yolo-server:v0.4.0 --label sedna=examples --no-cache ..