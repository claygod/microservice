package metrics

// Microservice
// Metrics
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	keyTitle = "route"
	keyCode  = "code"
)

type Metrics struct {
	registry *prometheus.Registry
	kits     *kits
	handler  http.Handler
}

func New() (*Metrics, error) {
	reg := prometheus.NewRegistry()

	m := &Metrics{
		registry: reg,
		kits:     getKits(),
		handler:  promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	}

	m.registry.MustRegister(
		collectors.NewGoCollector(),
	)

	if err := m.regKits(m.kits); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Metrics) Request(link string, duration time.Duration) {
	m.kits.request.cnt.With(prometheus.Labels{keyTitle: link}).Inc()

	m.kits.request.vec.
		With(prometheus.Labels{keyTitle: link}).
		Observe(float64(duration) / float64(time.Second))

	m.kits.request.his.
		With(prometheus.Labels{keyTitle: link}).
		Observe(float64(duration) / float64(time.Second))
}

func (m *Metrics) ResponceCode(respCode int) {
	m.kits.code.cnt.With(prometheus.Labels{keyTitle: strconv.Itoa(respCode)}).Inc()
}

func (m *Metrics) Handler() http.Handler {
	return m.handler
}

type kits struct {
	request *kit
	code    *kit
}

type kit struct {
	cnt *prometheus.CounterVec
	vec *prometheus.SummaryVec
	his *prometheus.HistogramVec
}

func (m *Metrics) regKits(k *kits) error {
	if err := m.regKit(k.request); err != nil {
		return err
	}

	if err := m.regKit(k.code); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) regKit(k *kit) error {
	if k.cnt != nil {
		if err := m.registry.Register(k.cnt); err != nil {
			return err
		}
	}

	if k.vec != nil {
		if err := m.registry.Register(k.vec); err != nil {
			return err
		}
	}

	if k.his != nil {
		if err := m.registry.Register(k.his); err != nil {
			return err
		}
	}

	return nil
}

func getKits() *kits {
	req := &kit{
		cnt: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "req_counter",
				Help: "Requests (counter).",
			}, []string{keyTitle},
		),
		vec: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:       "req_summary",
				Help:       "Requests (summary).",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			}, []string{keyTitle},
		),
		his: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "req_histogram",
			Help:    "Requests (histogram).",
			Buckets: []float64{1.0, 5.0, 10.0},
		}, []string{keyTitle}),
	}

	code := &kit{
		cnt: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "response_counter_code",
				Help: "Code Response Counter (counter).",
			}, []string{keyTitle},
		),
	}

	return &kits{
		request: req,
		code:    code,
	}
}
