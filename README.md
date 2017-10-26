# Go Workflow
This repository contains workflow microservice application written in Go.

This application contains 3 commands:
- migrate (workflow DB Schema migration)
- api (workflow api service)
- job (workflow job service)

API & Job command can be started on separated container

## Installation
- Download golang v1.9.1 here https://golang.org/dl/
- Install package from go get `go get github.com/kwri/go-workflow`

#### Compile

`go build main.go`

#### DB migration
##### Migrate
`go run main.go migrate migrate`
##### Rollback
`go run main.go migrate rollback`
##### Create migration script
`go run main.go migrate create your_script_name`

Why Go?
- Go is fast
- Go is Open source programming language
- Go is programming language that created for concurrency programming
- Go has strong standard library
- Go compiled to native language and can run on machine without JVM

## About each of the Services

- [About Rule Service](https://github.com/KWRI/go-workflow/wiki/About-Rule-Service)