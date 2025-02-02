package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	tr "go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
)

const jaegerURL = "http://localhost:14268/api/traces"

var tracer tr.Tracer

func main() {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("client"),
		)),
	)
	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)

	tracer = tp.Tracer("client tracer")

	http.HandleFunc("/tracing", exampleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
