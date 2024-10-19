package bar

// Microservice
// Bar gateway
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/claygod/microservice/domain"
)

const (
	secHelthBorder = 5
	emptyString    = ""
)

/*
BarGate - mock gateway
*/
type BarGate struct {
	hasp   domain.StartStopInterface
	logger *slog.Logger
	config *Config
}

func New(ss domain.StartStopInterface, lg *slog.Logger, cnf *Config) *BarGate {
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

	case barID == emptyString:
		b.logger.Warn(fmt.Sprintf("%s:length of id `%s` is zero", b.config.Title, barID))

		return nil, nil

	case len(barID) > b.config.MaxIDLenght:
		b.logger.Warn(fmt.Sprintf("%s:length of id `%s` is greater than %d", b.config.Title, barID, b.config.MaxIDLenght))

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
	if time.Now().Second() < secHelthBorder { // first 5 second error for example!!
		return http.StatusServiceUnavailable
	}

	return http.StatusOK
}

type Config struct {
	Title       string `toml:"gate_title" yaml:"gate_title" env:"Gate_title"`
	Prefix      string `toml:"gate_prefix" yaml:"gate_prefix" env:"GATE_PREFIX"`
	MaxIDLenght int    `toml:"gate_max_id_length" yaml:"gate_max_id_length" env:"GATE_MAX_ID_LENGTH"`
}
