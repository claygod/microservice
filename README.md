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
- `/healthz/ready` for SRE
- `/healthz` for kubernetes
- `/readyness` for kubernetes
- `/metrics` prometheus metrics
- `/piblic/v1/bar/:key` public route (example)

### Using

Build and run main.go

Example requests:

- localhost:8080/piblic/v1/bar/one -> {"Data":"three"}
- localhost:8080/piblic/v1/bar/secret -> response 404
- localhost:8080/piblic/v1/bar/looonnngggggkkkeeyyyyyyy -> response 404
- localhost:8080/healthz/ready -> minute first 5 sec - 503 after 200 (for example!)
- localhost:8080/healthz -> minute first 5 sec - 503 after 200 (for example!)
- localhost:8080/readyness -> response 200

### Environment

Add to ENV `export GATE_IN_TITLE=Yo-ho-ho!`
ang open in browser `http://localhost:8080/`

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

The default configuration file:
- `config/config.toml`

Specify in the command line another file:
- `yourservice -config other.toml`

## Dependencies

	github.com/claygod/tools v0.0.0-20211122181936-bab1329a2e3d
	github.com/dsbasko/go-cfg v1.2.0
	github.com/google/uuid v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pborman/getopt v1.1.0
	github.com/prometheus/client_golang v1.11.0

	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/caarlos0/env/v10 v10.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

## ToDo

- [x] Use environment variables in configuration
- [x] Add support for metrics
- [ ] Input validate
- [ ] Use protocol gRPC

## Conclusion

Microservice does not claim the laurels of the only true solution, but on occasion, I hope, will help you create your own micro-architecture of the service, becoming the prototype for future applications.

## Give us a star!

If you like or are using this project to learn or start your solution, please give it a star. Thank you!

## License

GNU GENERAL PUBLIC LICENSE Version 3

### Copyright © 2017-2024 Eduard Sesigin. All rights reserved. Contacts: claygod@yandex.ru
