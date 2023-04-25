#!/bin/bash

if ! sqlc version &> /dev/null
then
  echo "sqlc not found, installing it"
  $(go install github.com/kyleconroy/sqlc/cmd/sqlc@edb956047a10f8f928cdb8a131d18133c95ce809 && cd db && sqlc generate)
fi

if ! migrate -version &> /dev/null
then
  echo "migrate not found, installing it"
  go install --tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi

$(cd db && sqlc generate)
