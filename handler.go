// SPDX-License-Identifier: MIT

package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(opts ...func(*promhttp.HandlerOpts)) http.Handler {
	return HandlerFor(DefaultRegisterer, DefaultGatherer, opts...)
}

func HandlerFor(registerer prometheus.Registerer, gatherer prometheus.Gatherer, opts ...func(*promhttp.HandlerOpts)) http.Handler {
	opt := promhttp.HandlerOpts{}
	for _, o := range opts {
		o(&opt)
	}

	return promhttp.InstrumentMetricHandler(
		registerer, promhttp.HandlerFor(gatherer, opt),
	)
}
