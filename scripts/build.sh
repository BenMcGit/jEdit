#!/bin/bash

# get latest dependencies
go get -v

# build the module. This produces the jedit binary.
go build
