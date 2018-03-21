#!/usr/bin/env bash

echo - tar -
go run main.go test.tar

echo
echo ------------------
echo

echo - zip -
go run main.go test.zip
