#!/bin/sh
echo you may need to go get github.com/axw/gocov/gocov
gocov test ./... | gocov report