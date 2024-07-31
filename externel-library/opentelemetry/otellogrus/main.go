package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/bridges/otellogrus"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/log/noop"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	CONN      = "192.168.20.202:55680"
	HTTP_CONN = "http://192.168.20.202:55681" + "/v1/logs"
)

func main() {
	exam02()
	// exam01()
}

func exam02() {
	ctx := context.Background()

	// Create resource.
	res, err := newResource()
	//res, err := newResourceWithContext(ctx)
	if err != nil {
		panic(err)
	}

	// Create a logger provider.
	// You can pass this instance directly when creating bridges.
	loggerProvider, err := newLoggerProvider(ctx, res)
	if err != nil {
		panic(err)
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		if err := loggerProvider.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	// Register as global logger provider so that it can be accessed global.LoggerProvider.
	// Most log bridges use the global logger provider as default.
	// If the global logger provider is not set then a no-op implementation
	// is used, which fails to generate data.
	global.SetLoggerProvider(loggerProvider)

	// Create an *otellogrus.Hook and use it in your application.
	hook := otellogrus.NewHook("my-test", otellogrus.WithLoggerProvider(global.GetLoggerProvider()))

	// Set the newly created hook as a global logrus hook
	logrus.AddHook(hook)

	logrus.Debug("exam02: debug")
	logrus.Info("exam02: info")
	logrus.WithContext(ctx).Debug("exam02: context debug 1")
	logrus.WithContext(ctx).Info("exam02: context info 1")
	logrus.WithContext(ctx).Info("exam02: context info 2")
	logrus.WithContext(ctx).Info("exam02: context info 3")
	logrus.WithContext(ctx).Info("exam02: context info 4")
	logrus.WithContext(ctx).Info("exam02: context info 5")
}

func newResourceWithContext(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("provider"),
			semconv.ServiceVersion("0.0.0"),
		),
	)
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			//	semconv.SchemaURL,
			"https://opentelemetry.io/schemas/1.25.0",
			semconv.ServiceName("provider"),
			semconv.ServiceVersion("0.1.0"),
		))
}

func newLoggerProvider(ctx context.Context, res *resource.Resource) (*log.LoggerProvider, error) {
	// GRPC
	//ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//defer cancel()
	//conn, err := grpc.DialContext(ctx, CONN,
	//	// Note the use of insecure transport here. TLS is recommended in production.
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithBlock(),
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	//}
	//exporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))

	// HTTP
	exporter, err := otlploghttp.New(ctx, otlploghttp.WithEndpointURL(HTTP_CONN))
	if err != nil {
		return nil, err
	}
	processor := log.NewBatchProcessor(exporter)
	provider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(processor),
	)
	return provider, nil
}

func exam01() {
	// Use a working LoggerProvider implementation instead e.g. using go.opentelemetry.io/otel/sdk/log.
	provider := noop.NewLoggerProvider()
	//ctx := context.Background()
	//provider, err := collector(ctx, "hslog", "192.168.20.202:55680")
	//if err != nil {
	//	panic(err)
	//}

	// Create an *otellogrus.Hook and use it in your application.
	//hook := otellogrus.NewHook("my/pkg/name", otellogrus.WithLoggerProvider(provider))
	hook := otellogrus.NewHook("my/pkg/name", otellogrus.WithLoggerProvider(provider))

	// Set the newly created hook as a global logrus hook
	logrus.AddHook(hook)

	logrus.Debug("debug")
	logrus.Info("info")

	//logrus.WithContext(ctx).Info("ctx info")

}

// func collector(ctx context.Context, serviceName, collectorGrpcIp string) (func(context.Context) error, error) {
func collector(ctx context.Context, serviceName, collectorGrpcIp string) (*log.LoggerProvider, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// If the OpenTelemetry Collector is running on a local cluster (minikube or
	// microk8s), it should be accessible through the NodePort service at the
	// `localhost:30080` endpoint. Otherwise, replace `localhost` with the
	// endpoint of your cluster. If you run e app inside k8s, then you can
	// probably connect directly to the service through dns.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, collectorGrpcIp,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := log.NewBatchProcessor(logExporter)
	logProvider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(bsp),
	)

	return logProvider, nil

}
