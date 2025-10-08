// SPDX-License-Identifier: MIT

package exemplar

import (
	"context"

	"github.com/cdnnow-pro/go-tracer"
	"github.com/prometheus/client_golang/prometheus"
)

// FromContext makes a label to use as exemplar using the TraceID from SpanContext.
func FromContext(ctx context.Context) prometheus.Labels {
	if span := tracer.SpanFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{"traceID": span.TraceId()}
	}
	return nil
}
