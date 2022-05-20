# Quote API

[![License](https://img.shields.io/github/license/brianlusina/quote-api)](https://github.com/brianlusina/quote-api/blob/main/LICENSE)
[![Version](https://img.shields.io/github/v/release/brianlusina/quote-api?color=%235351FB&label=version)](https://github.com/brianlusina/quote-api/releases)
[![Tests](https://github.com/BrianLusina/quote-api/actions/workflows/tests.yml/badge.svg)](https://github.com/BrianLusina/quote-api/actions/workflows/tests.yml)
[![Lint](https://github.com/BrianLusina/quote-api/actions/workflows/lint.yml/badge.svg)](https://github.com/BrianLusina/quote-api/actions/workflows/lint.yml)
[![Build](https://github.com/BrianLusina/quote-api/actions/workflows/build_app.yml/badge.svg)](https://github.com/BrianLusina/quote-api/actions/workflows/build_app.yml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/3c7f8e37c31646a5ae7b17cf1682551b)](https://www.codacy.com/gh/BrianLusina/quote-api/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=BrianLusina/quote-api&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/BrianLusina/quote-api/branch/main/graph/badge.svg?token=RNg0UoESug)](https://codecov.io/gh/BrianLusina/quote-api)
[![Go](https://img.shields.io/badge/Go-1.18-blue.svg)](https://go.dev/)

Simple API to fetch quotes and add quotes.

## Requirements

A couple of things that you will need to set up the API up and running.

1. [Go](https://golang.org/doc/install)

   You will need to install the Go 1.18 in order to run the application.

2. [Docker](https://docker.com) and [Docker Compose](https://docs.docker.com/compose/install/)

    You will need to install Docker & docker compose in order to run services the application uses.

## Installation

Installation of dependencies can be done by running the following command:

``` bash
make install
```

> This install dependencies.

## Running the application

Before running the application, first setup environment variables.

``` bash
cp .env.sample .env
```

> sets up environment variables from .env.sample file. Set these environment variables approprately. The db configuration settings have been set to reasonable defaults.

now run the services specified in docker-compose.yml file.

``` bash
docker-compose up
```

Now we can run the application with the following command:

``` bash
make run
# or
go run app/cmd/main.go
```

> This will run the application on port 8080.

## Testing the application

Testing the application can be done by running the following command:

``` bash
make test
```

> Runs the tests.

Running test coverage can be done with:

``` bash
make test-coverage
```

## Linting application

Applicationg linting can be done by first setting up golangci-lint:

``` bash
make setup-linting
```

> This installs the [golangci-lint](https://github.com/golangci/golangci-lint) tool in the [bin directory](./bin).

Now, linting can be run with the below command:

``` bash
make lint
```

> Runs linting.

Additionally, linting can be done on the [Dockerfile](./Dockerfile) with [hadolint](https://github.com/hadolint/hadolint).

``` bash
make lint-docker
```

> This will run linting on the Dockerfile.

This uses [Docker](https://docker.com) to run the linting.
