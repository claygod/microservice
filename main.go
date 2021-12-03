package main

// Microservice
// Main
// This library provides a simple framework of microservice, which includes
// a configurator, a metrics, and of course the handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/claygod/microservice/app"
	"github.com/pborman/getopt"
)

const (
	defaultConfig               = "./config/config.toml"
	defaultUseEnv               = "false"
	envFieldsTag                = "env"
	shutdownTime  time.Duration = 1 * time.Minute // shutdown time limit
)

// Main
func main() {
	shutdown := make(chan bool)

	params := getCommandLineParameters()

	cnf, err := loadConfig(*params.ConfigPath)
	if err != nil {
		panic(err)
	}

	if params.EnvEnable != nil && *params.EnvEnable != defaultUseEnv {
		cnf.LoadEnvironment(envFieldsTag)
	}

	cmd, err := app.New(cnf)
	if err != nil {
		panic(err)
	}

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	gracefulStop(shutdown)
	<-shutdown

	go aggressiveStop()

	if err := cmd.Stop(); err != nil {
		panic(err)
	}
}

func getCommandLineParameters() *commandLineParameters {
	params := &commandLineParameters{
		ConfigPath: getopt.StringLong("config", 'c', defaultConfig, "Path to config file"),
		EnvEnable:  getopt.StringLong("env", 'e', defaultUseEnv, "Use environment variables in configuration "),
	}

	getopt.Parse()

	return params
}

type commandLineParameters struct {
	ConfigPath *string
	// PortHTTP   *string
	EnvEnable *string
}

func loadConfig(path string) (*app.Config, error) {
	var config app.Config

	_, err := toml.DecodeFile(path, &config)

	return &config, err
}

func gracefulStop(shutdown chan bool) {
	var gracefulStop = make(chan os.Signal)

	signal.Notify(gracefulStop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-gracefulStop
		// fmt.Printf("caught sig: %+v", sig)
		shutdown <- true
	}()
}

func aggressiveStop() {
	ticker := time.NewTicker(shutdownTime)

	<-ticker.C

	fmt.Println("The web application is aggressive stop")
	os.Exit(0)
}
