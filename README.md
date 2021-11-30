# Microservice

[![API documentation](https://godoc.org/github.com/claygod/microservice?status.svg)](https://godoc.org/github.com/claygod/microservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/microservice)](https://goreportcard.com/report/github.com/claygod/microservice)

The framework for the creation of microservices, written in Golang. 
This package is implemented using clean architecture principles
A good article on implementing these principles in Golang:
http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

## Endpoints

Code in `/services/gateways/gatein/gatein.go`

- `/` welcome handler
- `/health_check` for SRE
- `/readyness` for kubernetes
- `/piblic/v1/bar/:key` public route (example)

### Using

Build and run main.go

Example requests:

- localhost:8080/piblic/v1/bar/one
- localhost:8080/piblic/v1/bar/123
- localhost:8080/piblic/v1/bar/secret
- localhost:8080/piblic/v1/bar/looonnngggggkkkeeyyyyyyy

## Clean architecture

#### Entity

Path */domain*

#### Usecases

Path */usecases*

#### Interfaces

Path */service*

#### Infrastructure

Path */app* , */config* and core

## Config

The default configuration file:
- `config/config.toml`

Specify in the command line another file:
- `yourservice -confile other.toml`

## Dependencies

- github.com/BurntSushi/toml
- github.com/claygod/tools
- github.com/google/uuid
- github.com/julienschmidt/httprouter
- github.com/pborman/getopt
- github.com/sirupsen/logrus

## Conclusion

Microservice does not claim the laurels of the only true solution, but on occasion, I hope, will help you create your own micro-architecture of the service, becoming the prototype for future applications.

Copyright © 2017-2021 Eduard Sesigin. All rights reserved. Contacts: claygod@yandex.ru
