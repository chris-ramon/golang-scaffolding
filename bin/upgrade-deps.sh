#!/bin/bash

# Lists go dependencies along their latest version available.
go list -m -u all

# Updates all the go dependencies.
go get -u ./...

# Synchronizes go dependencies files.
go mod tidy
