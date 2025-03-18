// SPDX-License-Identifier: MIT

package exemplar

import (
	"context"

	"github.com/cdnnow-pro/tracer-go"
	"github.com/prometheus/client_golang/prometheus"
)

// ExemplarFromContext makes a label to use as exemplar using the TraceID from SpanContext.
func ExemplarFromContext(ctx context.Context) prometheus.Labels {
	if span := tracer.SpanFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{"traceID": span.TraceId()}
	}
	return nil
}
