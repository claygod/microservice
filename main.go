package main

// Microservice
// Main
// This library provides a simple framework of microservice, which includes
// a configurator, a metrics, and of course the handler
// Copyright © 2016-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/claygod/microservice/app"
	gocfg "github.com/dsbasko/go-cfg"
	"github.com/pborman/getopt"
)

const (
	defaultConfig               = "./config/config.toml" // alternative "./config/config.yml"
	shutdownTime  time.Duration = 1 * time.Minute        // shutdown time limit
)

func main() {
	shutdown := make(chan bool)

	params := getCommandLineParameters()

	cnf := loadConfigWithEnv(*params.ConfigPath)

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
	}

	getopt.Parse()

	return params
}

type commandLineParameters struct {
	ConfigPath *string
	// PortHTTP   *string
	EnvEnable *string
}

func loadConfigWithEnv(path string) *app.Config {
	var config app.Config

	gocfg.MustReadFile(path, &config)
	gocfg.MustReadEnv(&config)

	return &config
}

func gracefulStop(shutdown chan bool) {
	gracefulStop := make(chan os.Signal, 1)

	signal.Notify(gracefulStop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-gracefulStop
		shutdown <- true
	}()
}

func aggressiveStop() {
	ticker := time.NewTicker(shutdownTime)

	<-ticker.C

	fmt.Println("The web application is aggressive stop")
	os.Exit(0)
}
