// SPDX-License-Identifier: MIT

package metrics

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	_envNamespace = "METRICS_NAMESPACE"
	_envSubsystem = "METRICS_SUBSYSTEM"
)

var (
	Namespace string
	Subsystem string
)

func init() {
	if v, ok := os.LookupEnv(_envNamespace); ok {
		Namespace = v
	}
	if v, ok := os.LookupEnv(_envSubsystem); ok {
		Subsystem = v
	}
}

func NewCounterFor(r prometheus.Registerer, opts prometheus.CounterOpts) prometheus.Counter {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewCounter(opts)
}

func NewCounterVecFor(r prometheus.Registerer, opts prometheus.CounterOpts, labelValues []string) *prometheus.CounterVec {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewCounterVec(opts, labelValues)
}

func NewGaugeFor(r prometheus.Registerer, opts prometheus.GaugeOpts) prometheus.Gauge {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewGauge(opts)
}

func NewGaugeVecFor(r prometheus.Registerer, opts prometheus.GaugeOpts, labelValues []string) *prometheus.GaugeVec {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewGaugeVec(opts, labelValues)
}

func NewSummaryFor(r prometheus.Registerer, opts prometheus.SummaryOpts) prometheus.Summary {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewSummary(opts)
}

func NewSummaryVecFor(r prometheus.Registerer, opts prometheus.SummaryOpts, labelValues []string) *prometheus.SummaryVec {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewSummaryVec(opts, labelValues)
}

func NewHistogramFor(r prometheus.Registerer, opts prometheus.HistogramOpts) prometheus.Histogram {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewHistogram(opts)
}

func NewHistogramVecFor(r prometheus.Registerer, opts prometheus.HistogramOpts, labelValues []string) *prometheus.HistogramVec {
	if opts.Namespace == "" {
		opts.Namespace = Namespace
	}
	if opts.Subsystem == "" {
		opts.Subsystem = Subsystem
	}
	return promauto.With(r).NewHistogramVec(opts, labelValues)
}
