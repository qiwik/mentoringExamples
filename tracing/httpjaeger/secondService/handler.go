package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	propagator := otel.GetTextMapPropagator()
	pctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	traceID := r.Header.Get("trace-id")
	spanID := r.Header.Get("span-id")

	traceid, _ := trace.TraceIDFromHex(traceID)
	spanid, _ := trace.SpanIDFromHex(spanID)

	spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceid,
		SpanID:     spanid,
		TraceFlags: trace.FlagsSampled,
		Remote:     true,
	})

	sct := trace.ContextWithRemoteSpanContext(pctx, spanCtx)

	_, span := tracer.Start(sct, "server span")
	defer span.End()

	time.Sleep(2 * time.Second)

	return
}
