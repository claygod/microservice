# Microservice

[![API documentation](https://godoc.org/github.com/claygod/microservice?status.svg)](https://godoc.org/github.com/claygod/temp/microservice-doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/microservice)](https://goreportcard.com/report/github.com/claygod/microservice)

The framework for the creation of microservices, written in Golang.

Architecture microservice includes a handle, a tuner (configuration), place a couple of demonstration middleware and storage for them. All works is very simple: in the application configuration is loaded in view of the configuration file and command line environment. Created with the middleware *storage* and corresponding queues formed at the desired Route. Then run the server and request the application fulfills the desired queue.


## Tuner

The important point for microservice configuration is loaded. I tried to make this as flexible functionality. The default configuration is loaded from the specified file (`config.toml`). Address configuration file can be changed from the command line like this: `yourservice -confile config.toml` Thus, you can create several different configuration files and run microservice with one of the configurations of your choice.

Because the configuration file is not only, but also the environment variables, and command-line parameters are activated microservice, specify the order and priority configuration. The lowest priority is the configuration of the file. If the operating system environment variables are set up, they have a higher priority than variables from the configuration file. Command line parameters are most important priority. To change a parameter in the command line you need to specify its name in the form of a section name and parameter name (with capital letters!). Here is an example to change the port: `yourservice -Main/Port: 85`.

A little more about the configuration: in addition to the option with the introduction of changes in the structure prepared in advance, it could be used to import data option in ordinary *map*, and then safely use the values ​​for the key. This method has the undoubted advantage - no need to add data to the configuration file is duplicated in the configuration structure. Those. everything is easier and faster. The downside is that the wrong directions key errors transferred to compile at run time, and besides, there *IDE* no longer will we do tips that way itself whelp well to protect against typos and as a consequence - errors.

## Storage

In fact, the middleware storage facilities *Storage* not necessarily for use in microservice, and I'm not trying to implement anything from the fact that is often associated with the *DI*. Here rather subjective usability create and initialize structures in one place. Moreover, thanks to store *IDE* kindly name suggests their structures and methods exported.

## Middleware

In order not to clutter the space handlers, in the microservice use of the service provided by the functional style *middleware*. Each such service predyavlyat minimum requirements: it must take as arguments `http.ResponseWriter, http.Request` and return them. The distribution shows an example of connecting the metrics before and after the handler to lock request processing time (*duration*).

Depending on what tasks are microservices and the environment in which he does it, almost certainly you will need other services, such as validation. Look for them to connect to the handler as *middleware*. And note that the demonstration of the metric and the session in a separate subdirectory to emphasize the distance between the microservice and used in it *middleware*.

## Queue

This entity is just as easy as everything else. It simply keeps a list of functions (handler and midlvar), and when necessary, launches them for execution by Run. In this case, if some of the methods return `nil` instead of `http.Request`, turn stops the execution (eg it could do validator well or Access Controller).

## Handler

In the framework, *handler* contains very little code and is engaged in general that creates and returns the run queue *queue*. However, it is the methods *handler* and are the handler that will process the incoming request, the rest is essentially a microservice kit. If so, then why not just do a bunch of autonomous functions, which will cause the router? For good reason: thanks to the association by the structure handlers methods will now be able to access (if necessary) to the fields of the structure (general application context). In fact, very convenient for each public method to create a separate file, such as *handle_hello_word.go*, which, incidentally, does not interfere with how to organize everything you want in other ways.

## The algorithm works with a microservice

### Create a new Middleware

For example, you want to add a validator as *middleware*. In this case, the first in the configuration file *config.toml* section and add in it the necessary parameters. For example like this:
```toml
[Validation]
Number	= "integer"
```
Now create a corresponding structure (eg *ConfValidation*) in the file *config.go  and add it in the same place in the structure *Config*. Next to *storage* create a method for creating a validator, and then it can be used in the formation of a queue in *main.go*

### Creating a new handler

To do this, you just need to create a new public method in the *handler*, which takes input *http.ResponseWriter, http.Request* and returns them. Look created to demonstrate the method *HelloWorld*.

## Perfomance

For a general understanding of what is the speed of microservice using the proposed architecture will be high, and bring the benchmark results obtained by me on my computer IntelCore i3-2120:

- BenchmarkMain-4           	 5000000	       261 ns/op
- BenchmarkMainParallel-4   	10000000	       137 ns/op

## Dependencies

- Logger	https://github.com/Sirupsen/logrus
- Session	https://github.com/gorilla/sessions
- Route	https://github.com/claygod/Bxog
- Config	https://github.com/BurntSushi/toml

Any of these libraries can be replaced or supplemented, in this case, they are likely designed to show which way to develop their own microservices. You might also be useful to connect *logstash* and *influxdb*.

## Conclusion

Microservice Library does not claim the laurels of the only true solution, but on occasion, I hope, will help you create your own micro-architecture of the service, becoming the prototype for future applications.
