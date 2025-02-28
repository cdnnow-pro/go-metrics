package metrics

import (
	"os"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	_envNamespace = "METRICS_NAMESPACE"

	_envWithoutProcessMetrics   = "METRICS_WITHOUT_PROCESS"
	_envWithoutCPUMetrics       = "METRICS_WITHOUT_CPU"
	_envWithoutGCMetrics        = "METRICS_WITHOUT_GC"
	_envWithoutMemoryMetrics    = "METRICS_WITHOUT_MEMORY"
	_envWithoutSchedulerMetrics = "METRICS_WITHOUT_SCHEDULER"
	_envWithoutSyncMetrics      = "METRICS_WITHOUT_SYNC"
)

func init() {
	if v, ok := os.LookupEnv(_envNamespace); ok {
		namespace = v
	}

	if withProcessMetrics() {
		DefaultRegisterer.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	var rules []collectors.GoRuntimeMetricsRule
	if withCPUMetrics() {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/cpu/classes/.*`)})
	}
	if withGCMetrics() {
		rules = append(rules, collectors.MetricsGC)
	}
	if withMemoryMetrics() {
		rules = append(rules, collectors.MetricsMemory)
	}
	if withSchedulerMetrics() {
		rules = append(rules, collectors.MetricsScheduler)
	}
	if withSyncMetrics() {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/sync/.*`)})
	}
	if len(rules) > 0 {
		DefaultRegisterer.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(rules...)))
	}
}

func withProcessMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutProcessMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withCPUMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutCPUMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withGCMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutGCMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withMemoryMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutMemoryMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withSchedulerMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutSchedulerMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withSyncMetrics() bool {
	if v, ok := os.LookupEnv(_envWithoutSyncMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}
