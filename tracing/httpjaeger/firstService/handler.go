package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net/http"
	"time"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	spanCtx, span := tracer.Start(r.Context(), "exampleHandler in client")
	defer span.End()

	req, err := http.NewRequest("GET", "http://localhost:8081/example", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("trace-id", span.SpanContext().TraceID().String())
	req.Header.Set("span-id", span.SpanContext().SpanID().String())

	p := otel.GetTextMapPropagator()
	p.Inject(spanCtx, propagation.HeaderCarrier(req.Header))

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}
