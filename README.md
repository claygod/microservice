# Microservice

[![API documentation](https://godoc.org/github.com/claygod/microservice?status.svg)](https://godoc.org/github.com/claygod/microservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/microservice)](https://goreportcard.com/report/github.com/claygod/microservice)

The framework for the creation of microservices, written in Golang. (note: http microservice)

Architecture microservice includes:

- handle
- tuner (configuration)
- middleware-style tools (for demo)

### Create a new Middleware

Use as a example *tools/metric.go* file. 

## Creating a new handler

To do this, you just need to create a new public method in the *handler*, which takes input *http.ResponseWriter, http.Request*. Look created to demonstrate the method *HelloWorld*.

## Perfomance

For a general understanding of what is the speed of microservice using the proposed architecture will be high, and bring the benchmark results obtained by me on my computer:

- BenchmarkMain-2            	10000000	       192 ns/op
- BenchmarkMainParallel-2   	10000000	       104 ns/op

## Tuner

The default configuration file:
- `config.toml`

Specify in the command line another file:
- `yourservice -confile other.toml`

To change the setting on the command line, you specify the section and parameter name (composed by a slash): 
- `yourservice -Main/Port 85`

Configuring priorities:
- command line (highest priority)
- environment
- configuration file

## Dependencies

- Logger	https://github.com/Sirupsen/logrus
- Route	https://github.com/claygod/Bxog
- Config	https://github.com/BurntSushi/toml

Any of these libraries can be replaced or supplemented, in this case, they are likely designed to show which way to develop their own microservices. You might also be useful to connect *logstash* and *influxdb*.

## Conclusion

Microservice Library does not claim the laurels of the only true solution, but on occasion, I hope, will help you create your own micro-architecture of the service, becoming the prototype for future applications.
