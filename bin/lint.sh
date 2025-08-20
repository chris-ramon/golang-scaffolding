#!/bin/bash

if ! golangci-lint version &> /dev/null
then
  echo "golangci-lint not found, installing it"
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

golangci-lint run ./...