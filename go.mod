module apps

go 1.19

require (
	go.opentelemetry.io/contrib/propagators/b3 v1.15.0
	go.opentelemetry.io/otel v1.15.0-rc.2
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.15.0-rc.2
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v0.0.0
	go.opentelemetry.io/otel/sdk v1.15.0-rc.2
)

replace (
	go.opentelemetry.io/otel => ../opentelemetry-go-query
	go.opentelemetry.io/otel/exporters/otlp/internal/retry => ../opentelemetry-go-query/exporters/otlp/internal/retry
	go.opentelemetry.io/otel/exporters/otlp/otlptrace => ../opentelemetry-go-query/exporters/otlp/otlptrace
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp => ../opentelemetry-go-query/exporters/otlp/otlptrace/otlptracehttp
	go.opentelemetry.io/otel/sdk => ../opentelemetry-go-query/sdk
	go.opentelemetry.io/otel/semconv => ../opentelemetry-go-query/semconv
	go.opentelemetry.io/otel/trace => ../opentelemetry-go-query/trace
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.15.0-rc.2 // indirect
	go.opentelemetry.io/otel/trace v1.15.0-rc.2 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
