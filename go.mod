module github.com/open-telemetry/opentelemetry-lambda-extension

go 1.14

require (

	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter v0.20.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder v0.20.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/collector v0.20.1-0.20210218001603-48151d869607
	go.uber.org/zap v1.16.0
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/collector v0.14.0
	go.uber.org/zap v1.16.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ./extension/observer

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver => ./extension/observer/hostobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver => ./extension/observer/k8sobserver

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ./internal/k8sconfig
