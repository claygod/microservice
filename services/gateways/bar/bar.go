package bar

// Microservice
// Bar gateway
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"net/http"
	"time"

	"github.com/claygod/microservice/domain"
	"github.com/sirupsen/logrus"
)

/*
BarGate - mock gateway
*/
type BarGate struct {
	hasp   domain.StartStopInterface
	logger *logrus.Entry
	config *Config
}

func New(ss domain.StartStopInterface, lg *logrus.Entry, cnf *Config) *BarGate {
	return &BarGate{
		hasp:   ss,
		logger: lg,
		config: cnf,
	}
}

func (b *BarGate) GetBar(barID string) (*domain.Bar, error) {
	switch {
	case barID == "two":
		return &domain.Bar{Data: "three"}, nil

	case barID == "hide": // not found
		return nil, nil

	case len(barID) == 0:
		b.logger.Warn(fmt.Errorf("%s:length of id `%s` is zero", b.config.Title, barID))

		return nil, nil

	case len(barID) > *b.config.MaxIDLenght:
		b.logger.Warn(fmt.Errorf("%s:length of id `%s` is greater than %d", b.config.Title, barID, b.config.MaxIDLenght))

		return nil, nil

	default:
		return &domain.Bar{Data: fmt.Sprintf("%s:%s", b.config.Prefix, barID)}, nil

	}
}

func (b *BarGate) Start() error {
	if !b.hasp.Start() {
		return fmt.Errorf("%s:failed to start", b.config.Title)
	}

	return nil
}

func (b *BarGate) Stop() error {
	if !b.hasp.Stop() {
		return fmt.Errorf("%s:failed to stop", b.config.Title)
	}

	return nil
}

func (b *BarGate) CheckStatus() int {
	if time.Now().Second() < 20 {
		return http.StatusServiceUnavailable
	}

	return http.StatusOK
}

type Config struct {
	Title       string
	Prefix      string
	MaxIDLenght *int `env:"BAR_MAX_ID_LENGHT"`
}
