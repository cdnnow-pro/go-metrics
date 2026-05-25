package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/cdnnow-pro/go-metrics"
	"github.com/cdnnow-pro/go-metrics/exemplar"
	"github.com/cdnnow-pro/go-tracer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := prometheus.NewRegistry()

	var (
		test1 = metrics.NewCounterVecFor(r, prometheus.CounterOpts{
			Name: "test_total",
			Help: "Test1 help",
		}, []string{
			"label1",
			"label2",
		})
		test2 = metrics.NewHistogramVecFor(r, prometheus.HistogramOpts{
			Name:    "test_seconds",
			Help:    "Test2 help",
			Buckets: prometheus.DefBuckets,
		}, []string{
			"label3",
		})
	)

	http.Handle("/metrics", metrics.HandlerFor(r, r, func(opts *promhttp.HandlerOpts) {
		opts.EnableOpenMetrics = true
	}))
	go http.ListenAndServe(":8089", http.DefaultServeMux)

	ctx := context.Background()

	closer, err := tracer.Init(ctx, "metrics_test", "v1.0.0")
	if err != nil {
		panic(err)
	}
	defer closer(ctx)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()

	ctx, _ = tracer.StartSpan(ctx, "main")
	exemplarLabel := exemplar.FromContext(ctx)

	test1.With(prometheus.Labels{"label1": "A1", "label2": "B1"}).Inc()
	test1.With(prometheus.Labels{"label1": "A1", "label2": "B2"}).Inc()
	test1.With(prometheus.Labels{"label1": "A1", "label2": "B2"}).Inc()

	test1.With(prometheus.Labels{"label1": "A2", "label2": "B3"}).(prometheus.ExemplarAdder).AddWithExemplar(8, exemplarLabel)

	test2.With(prometheus.Labels{"label3": "N"}).(prometheus.ExemplarObserver).ObserveWithExemplar(0.425, exemplarLabel)

	<-ctx.Done()
}
