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

var (
	test1 = metrics.NewCounterVec("test_total", "Test1 help", []string{
		"label1",
		"label2",
	})
	test2 = metrics.NewHistogramVec("test_seconds", "Test2 help", metrics.ResponseTimeBuckets, []string{
		"label3",
	})
)

func main() {
	http.Handle("/metrics", metrics.Handler(func(opts *promhttp.HandlerOpts) {
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
	exemplarLabel := exemplar.ExemplarFromContext(ctx)

	test1.With(prometheus.Labels{"label1": "A1", "label2": "B1"}).Inc()
	test1.With(prometheus.Labels{"label1": "A1", "label2": "B2"}).Inc()
	test1.With(prometheus.Labels{"label1": "A1", "label2": "B2"}).Inc()

	test1.With(prometheus.Labels{"label1": "A2", "label2": "B3"}).(prometheus.ExemplarAdder).AddWithExemplar(8, exemplarLabel)

	test2.With(prometheus.Labels{"label3": "N"}).(prometheus.ExemplarObserver).ObserveWithExemplar(0.425, exemplarLabel)

	<-ctx.Done()
}
