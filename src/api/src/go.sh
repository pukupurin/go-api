#!/bin/sh
cd /root/api
go mod tidy
go run main.go >> /root/log/api.log