# Fibonacci Web Service (Go)

A simple web service that returns the first `n` Fibonacci numbers.

The service has one endpoint `/fibonacci` and always expects the `n` query param. This specifies how many fibonacci numbers you would like to generate.

## Prerequisites
Go >= 1.24.4

## Quick Start
1. Build and run the service:
```bash
make deploy
```
2. Use the API:
```bash
curl "http://localhost:8080/fibonacci?n=100&big=true"
```
## Usage
Run make to see available commands.
```bash
Usage:
  help          print this help message
  tidy          tidy modfiles and format .go files
  clean         clean build directory
  build         build binary
  run           runs the binary
  deploy        builds and runs the binary
  audit         run quality control checks
  test          run all tests
  test/cover    run all tests and display coverage
  test/bench    run all benchmarks
  upgradeable   list direct dependencies that have upgrades available
```

## Test
```bash
make audit
```
Make audit will run all test types.


## Build and Run on Local Machine
```bash
make deploy
```
Make deploy will build and run the binary.

## Build and Run in Docker Container
```bash
make docker
```
Make docker will build and run a docker container.
