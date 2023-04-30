package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"io"
	"net/http"
	"strconv"
)

var tracer = otel.Tracer("BDMLSys-app")

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithInsecure())
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	return exporter, nil
}

func newTracerProvider(exporter sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("Fibonacci")))

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
}

func Fibonacci(n int64, ctx context.Context) (int64, error) {
	ctx, span := tracer.Start(ctx, "Fibonacci")
	defer span.End()

	if n <= 1 {
		return n, nil
	}

	if n > 93 {
		span.AddEvent("request n too large")
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}

	var n2, n1 int64 = 0, 1
	for i := int64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

func calcFib(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "calcFib")
	defer span.End()

	span.AddEvent("calculating fibonacci number")

	reqN, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		span.SetAttributes(
			attribute.String("test1", "error"),
			attribute.Int("request_n", -1))
		return
	}

	span.SetAttributes(
		attribute.String("test1", "ok"),
		attribute.Int64("request_n", int64(reqN)))

	res, err := Fibonacci(int64(reqN), ctx)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	span.SetAttributes(
		attribute.Int64("final_result", res),
		attribute.Bool("calc_success", true))

	_, err = io.WriteString(w, strconv.FormatInt(res, 10))
	if err != nil {
		span.AddEvent("error writing response")
		return
	}
}

func refreshFilter(w http.ResponseWriter, r *http.Request) {
	taf := otel.GetTraceAttributeFilter()
	println("Start refreshing filter")
	err := taf.HandleRequest(r)
	if err != nil {
		println("Error refreshing filter")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	println("Finish refreshing filter")
	_, err = io.WriteString(w, "ok")
	if err != nil {
		return
	}
}

func main() {
	ctx := context.Background()
	exporter, err := newExporter(ctx)
	if err != nil {
		panic(err)
	}

	tp := newTracerProvider(exporter)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	p := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	otel.SetTextMapPropagator(p)

	http.HandleFunc("/fib", calcFib)
	http.HandleFunc("/otel", refreshFilter)

	_ = http.ListenAndServe(":1233", nil)
}
