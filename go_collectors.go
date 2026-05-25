// SPDX-License-Identifier: MIT

package metrics

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	withoutProcess uint = 1 << iota
	withoutCPU
	withoutGC
	withoutMemory
	withoutScheduler
	withoutSync
	withMemStats
)

type flags uint

func (ff flags) noFlag(f uint) bool {
	return uint(ff)&f == 0
}

func (ff flags) hasFlag(f uint) bool {
	return uint(ff)&f == 1
}

type GoCollectorsOption func(*flags)

func WithoutProcess() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutProcess)
	}
}
func WithoutCPU() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutCPU)
	}
}
func WithoutGC() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutGC)
	}
}
func WithoutMemory() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutMemory)
	}
}
func WithoutScheduler() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutScheduler)
	}
}
func WithoutSync() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withoutSync)
	}
}
func WithMemStats() GoCollectorsOption {
	return func(o *flags) {
		*o = flags(uint(*o) | withMemStats)
	}
}

func MustRegisterGoCollectors(r prometheus.Registerer, opts ...GoCollectorsOption) {
	var ff flags
	for _, opt := range opts {
		opt(&ff)
	}

	if ff.noFlag(withoutProcess) {
		r.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	var rules []collectors.GoRuntimeMetricsRule
	if ff.noFlag(withoutCPU) {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/cpu/classes/.*`)})
	}
	if ff.noFlag(withoutGC) {
		rules = append(rules, collectors.MetricsGC)
	}
	if ff.noFlag(withoutMemory) {
		rules = append(rules, collectors.MetricsMemory)
	}
	if ff.noFlag(withoutScheduler) {
		rules = append(rules, collectors.MetricsScheduler)
	}
	if ff.noFlag(withoutSync) {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/sync/.*`)})
	}

	switch {
	case len(rules) == 0 && ff.hasFlag(withMemStats):
		r.MustRegister(collectors.NewGoCollector(
			collectors.WithGoCollectorMemStatsMetricsDisabled(),
		))
	case ff.hasFlag(withMemStats):
		r.MustRegister(collectors.NewGoCollector(
			collectors.WithGoCollectorMemStatsMetricsDisabled(),
			collectors.WithGoCollectorRuntimeMetrics(rules...),
		))
	default:
		r.MustRegister(collectors.NewGoCollector(
			collectors.WithGoCollectorRuntimeMetrics(rules...),
		))
	}
}
