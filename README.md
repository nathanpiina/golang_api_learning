# Project Name

## Description

This project corresponds to the creation of a basic API using Golang as a language, the structure and objectives of the API are based on the "[Rinha de Backend - 2023](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/blob/main/INSTRUCOES.md)".

## How to Run

You'll need to have [docker](https://docs.docker.com/engine/install/ubuntu/), [docker compose](https://docs.docker.com/compose/cli-command/#install-on-linux) plugin, etc... installed.

* **Install dependencies**

```bash
$ make install
$ yarn install
$ npm i
$ ./gradlew clean build
$ ./mvnw clean install
$ go mod vendor
# write here an example output of this command
```

* **Running local dependencies** if you want to run the application using local dependencies with docker, run the command below before running the codebase

```bash
$ make dependencies/up
$ docker compose up
# write here an example output of this command
```

* **Run app** in dev mode (watch changes and reload)

```bash
$ make dev/api
$ npm start
$ yarn dev
$ ./mvnw spring-boot:run
$ ./gradlew bootRun
# write here an example output of this command
```

## How to Test

### Unit

Just run:

```bash
$ make test/unit
$ npm run test
$ ./gradlew test
$ ./mvnw test
$ go test ./...
# write here an example output of this command
```

### Integration

Just run:

```bash
$ make test/integration
$ npm run itest
$ ./gradlew integrationTest
$ ./mvnw test
$ go test ./...
# write here an example output of this command
```

### E2E

Just run

```bash
$ make test/e2e
$ npm run e2eTest
$ ./gradlew e2eTest
# write here an example output of this command
```

### Manual testing

Describe here what you can do to perform a manual test :D

```bash
$ curl -XPOST -d '{"id":1}' http://localhost:9000
# write here an example output of this command
```

### Lint

Just run:

```bash
$ make lint
$ npm run lint
$ ./gradlew spotlessApply
$ ./mvnw check
# write here an example output of this command
```

## How to Deploy

Describe here what you can do to deploy the application code into the existent environments, you can detail all your CI/CD here!

Eg:

Deployment is automated via github actions.

* When changes are pushed to `main` branch code goes to DEV env.
* When release is created code goes to PRD env.

## Built With

* Lang 1.18
* Docker
* etc
* Cloud functions
* Serverless Framework
* Kafka
* Redis
* Postgres
* Draw if you want!

## Application Checklist

## Other Monitoring/Observability

Add links to any custom Datadog dashboards, GCP functions, alert policies etc.
Document everything that is used to monitor this service, so everyone will know what is covered and where to look.

## Architecture

* Diagram

```bash
Put a PNG of your architecture here!
```

* Environments

| Env  | URL                     |
| ---- | ----------------------- |
| LOCAL| <http://localhost:9000>   |
