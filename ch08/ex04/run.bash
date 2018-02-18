#!/usr/bin/env bash

go run ./reverb2/reverb.go &
sleep 1
go run ./netcat3/netcat.go
