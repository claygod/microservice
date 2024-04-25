# Microservice

[![API documentation](https://godoc.org/github.com/claygod/microservice?status.svg)](https://godoc.org/github.com/claygod/microservice)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![forks](https://img.shields.io/github/forks/claygod/microservice)](https://github.com/claygod/microservice/network/members)
[![stars](https://img.shields.io/github/stars/claygod/microservice)](https://github.com/claygod/microservice/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/microservice)](https://goreportcard.com/report/github.com/claygod/microservice)

The framework for the creation of microservices, written in Golang. 
This package is implemented using clean architecture principles.
A good article on implementing these principles in Golang:
http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

## Endpoints

Code in `/services/gateways/gatein/gatein.go`

- `/` welcome handler
- `/health_check` for SRE
- `/readyness` for kubernetes
- `/metrics` prometheus metrics
- `/piblic/v1/bar/:key` public route (example)

### Using

Build and run main.go

Example requests:

- localhost:8080/piblic/v1/bar/one
- localhost:8080/piblic/v1/bar/123
- localhost:8080/piblic/v1/bar/secret -> response 404
- localhost:8080/piblic/v1/bar/looonnngggggkkkeeyyyyyyy

## Clean architecture

Distribution of architectural entities by layers

#### Entity

Path */domain*

#### Usecases

Path */usecases*

#### Interfaces

Path */service*

#### Infrastructure

Path */app* , */config* and core

## Config

Types of variables that support the use of environment variables
in the configuration (a pointer is required, or will be ignored!) :
- string
- float32, float64
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64

The default configuration file:
- `config/config.toml`

Specify in the command line another file:
- `yourservice -config other.toml`

Use environment variables in configuration:
- `yourservice -env true`

Command line xample:
- `foo -config stage.toml -env true`

Configuration with environment tag example:
```Golang
type Config struct {
	MaxIDLenght *int `env:"FOO_MAX_ID_LENGHT"`
}
```

## Dependencies

- github.com/BurntSushi/toml
- github.com/claygod/tools
- github.com/google/uuid
- github.com/julienschmidt/httprouter
- github.com/pborman/getopt
- github.com/prometheus/client_golang
- github.com/prometheus/tsdb
- github.com/sirupsen/logrus

## ToDo

- [x] Use environment variables in configuration
- [x] Add support for metrics
- [ ] Use protocol gRPC

## Conclusion

Microservice does not claim the laurels of the only true solution, but on occasion, I hope, will help you create your own micro-architecture of the service, becoming the prototype for future applications.

## Give us a star!

If you like or are using this project to learn or start your solution, please give it a star. Thank you!

## License

GNU GENERAL PUBLIC LICENSE Version 3

### Copyright © 2017-2024 Eduard Sesigin. All rights reserved. Contacts: claygod@yandex.ru
