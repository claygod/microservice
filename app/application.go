package app

// Microservice
// Application
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"log/slog"
)

type Application struct {
	logger    *slog.Logger
	listToRun []Item
}

func (a *Application) Start() error {
	err := a.exeStartList(a.listToRun)

	if err != nil {
		a.logger.Error(err.Error())
	} else {
		a.logger.Info("Application is started")
	}

	return err
}

func (a *Application) Stop() error {
	err := a.exeStopList(a.listToRun)

	if err != nil {
		a.logger.Error(err.Error())
	} else {
		a.logger.Info("Application is stopped")
	}

	return err
}

func (a *Application) exeStartList(list []Item) error {
	revertList := make([]Item, 0, len(list))

	for i := 0; i < len(list); i++ {
		if err := list[i].Start(); err != nil {
			return fmt.Errorf("%v: %w", err, a.exeStopList(revertList))
		}

		revertList = append(revertList, list[i])
	}

	return nil
}

func (a *Application) exeStopList(list []Item) error {
	var allErr error

	for i := len(list) - 1; i >= 0; i-- {
		if err := list[i].Stop(); err != nil {
			a.logger.Error(err.Error())
			allErr = fmt.Errorf("%v: %w", allErr, err)
		}
	}

	return allErr
}

type Item interface {
	Start() error
	Stop() error
}
