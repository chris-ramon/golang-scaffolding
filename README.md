# golang-scaffolding [![Tests](https://github.com/chris-ramon/golang-scaffolding/actions/workflows/tests.yml/badge.svg)](https://github.com/chris-ramon/golang-scaffolding/actions/workflows/tests.yml) [![codecov](https://codecov.io/gh/chris-ramon/golang-scaffolding/branch/main/graph/badge.svg?token=VUGFGVC37X)](https://codecov.io/gh/chris-ramon/golang-scaffolding)

A golang scaffolding for getting started new projects.

#### Getting Started

##### Replace Paths

```
ag -l 'chris-ramon/golang-scaffolding'| xargs sed -i 's|chris-ramon/golang-scaffolding|<your-org>/<your-repo>|g'
```

##### Create `.env` file

`APP_RSA`
```
cat app.rsa | base64 | tr -d '\n'|pbcopy
```

`APP_RSA_PUB`
```
cat app.rsa.pub | base64 | tr -d '\n'|pbcopy
```

##### Running

```
./bin/dev.sh
```

##### GraphQL Playground

[http://localhost:8080/graphql](http://localhost:8080/graphql)

##### DB Migrations

```bash
docker exec -it golang-scaffolding-app-1 bash
```

Up

```bash
migrate -database "postgres://admin:admin@db:5432/local?sslmode=disable" -path "./db/migrations" up 1
```

Down

```bash
migrate -database "postgres://admin:admin@db:5432/local?sslmode=disable" -path "./db/migrations" down 1
```


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
