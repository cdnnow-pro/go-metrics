// SPDX-License-Identifier: MIT

package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	namespace = ""
	registry  = prometheus.NewRegistry()

	DefaultRegisterer prometheus.Registerer = registry
	DefaultGatherer   prometheus.Gatherer   = registry

	DefaultSummaryObjectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001} //nolint:mnd

	ResponseTimeBuckets = []float64{0.001, 0.003, 0.007, 0.01, 0.015, 0.05, 0.1, 0.3, 0.6, 1, 3, 6, 9, 15, 30}
)

func Register(c prometheus.Collector) error {
	return DefaultRegisterer.Register(c)
}

func MustRegister(cs ...prometheus.Collector) {
	DefaultRegisterer.MustRegister(cs...)
}

func Unregister(c prometheus.Collector) bool {
	return DefaultRegisterer.Unregister(c)
}

func NewCounter(name, help string, r ...prometheus.Registerer) prometheus.Counter {
	return promauto.With(registerer(r)).NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewGauge(name, help string, r ...prometheus.Registerer) prometheus.Gauge {
	return promauto.With(registerer(r)).NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewHistogram(name, help string, buckets []float64, r ...prometheus.Registerer) prometheus.Histogram {
	return promauto.With(registerer(r)).NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	})
}

func NewSummary(name, help string, objectives map[float64]float64, r ...prometheus.Registerer) prometheus.Summary {
	return promauto.With(registerer(r)).NewSummary(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	})
}

func NewCounterVec(name, help string, labelValues []string, r ...prometheus.Registerer) *prometheus.CounterVec {
	return promauto.With(registerer(r)).NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelValues)
}

func NewGaugeVec(name, help string, labelValues []string, r ...prometheus.Registerer) *prometheus.GaugeVec {
	return promauto.With(registerer(r)).NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelValues)
}

func NewHistogramVec(name, help string, buckets []float64, labelValues []string, r ...prometheus.Registerer) *prometheus.HistogramVec {
	return promauto.With(registerer(r)).NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labelValues)
}

func NewSummaryVec(name, help string, objectives map[float64]float64, labelValues []string, r ...prometheus.Registerer) *prometheus.SummaryVec {
	return promauto.With(registerer(r)).NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	}, labelValues)
}

func registerer(rr []prometheus.Registerer) prometheus.Registerer {
	if len(rr) == 0 {
		return DefaultRegisterer
	}

	if len(rr) > 1 {
		log.Println("multiple registerer not supported")
	}
	return rr[0]
}
