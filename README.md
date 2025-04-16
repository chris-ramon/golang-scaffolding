# golang-scaffolding [![Tests](https://github.com/chris-ramon/golang-scaffolding/actions/workflows/tests.yml/badge.svg)](https://github.com/chris-ramon/golang-scaffolding/actions/workflows/tests.yml) [![codecov](https://codecov.io/gh/chris-ramon/golang-scaffolding/branch/main/graph/badge.svg?token=VUGFGVC37X)](https://codecov.io/gh/chris-ramon/golang-scaffolding)

A golang scaffolding for getting started new projects.

#### Getting Started

##### Replace Paths

```
ag -l 'chris-ramon/golang-scaffolding'|xargs sed -i 's/<your-org>/<your-repo>/g'
```

##### Running

```
docker compose up
```

##### GraphQL Playground

[http://localhost:8080/graphql](http://localhost:8080/graphql)


##### Features

Contains the following example domains:
- [x] Env Variables.
- [x] Config.
- [x] Auth.
- [x] JWT.
- [x] GraphQL.
- [x] PostgreSQL.
- [x] Type Safe SQL.
- [x] Docker Compose.
- [x] Live reload.
- [x] Admin.
- [x] Unit tests.
