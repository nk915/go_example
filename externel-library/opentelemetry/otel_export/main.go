// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Example using OTLP exporters + collector + third-party backends. For
// information about using the exporter, see:
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp?tab=doc#example-package-Insecure
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

// Initializes an OTLP exporter, and configures the corresponding trace and
// metric providers.
func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName("test-service"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// If the OpenTelemetry Collector is running on a local cluster (minikube or
	// microk8s), it should be accessible through the NodePort service at the
	// `localhost:30080` endpoint. Otherwise, replace `localhost` with the
	// endpoint of your cluster. If you run the app inside k8s, then you can
	// probably connect directly to the service through dns.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "192.168.20.110:4317",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

func main() {
	log.Printf("Waiting for connection...")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initProvider()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := otel.Tracer("test-tracer")

	// Attributes represent additional key-value descriptors that can be bound
	// to a metric observer or recorder.
	commonAttrs := []attribute.KeyValue{
		attribute.String("attrA", "chocolate"),
		attribute.String("attrB", "raspberry"),
		attribute.String("attrC", "vanilla"),
	}

	// work begins
	ctx, span := tracer.Start(
		ctx,
		"CollectorExporter-Example",
		trace.WithAttributes(commonAttrs...))

	// SpanContext 추출
	//sc := trace.SpanFromContext(ctx).SpanContext()
	sc := span.SpanContext()

	// SpanContext를 문자열로 변환
	scStr := sc.TraceID().String() + "-" + sc.SpanID().String()

	fmt.Println("trace: ", scStr)

	defer span.End()
	for i := 0; i < 10; i++ {
		_, iSpan := tracer.Start(ctx, fmt.Sprintf("Sample-%d", i))
		log.Printf("Doing really hard work (%d / 10)\n", i+1)

		<-time.After(time.Second)
		iSpan.End()
	}

	fmt.Println("trace: ", scStr)
	log.Printf("Done!")
	// select {}
	// log.Printf("Done!")
}

func trace2() {
	//	ctx := context.Background()
	//
	//	// Collector의 주소를 설정합니다.
	//	driver := exporter.NewDriver(
	//		exporter.WithInsecure(),
	//		exporter.WithEndpoint("localhost:55680"),
	//	)
	//
	//	exporter, err := exporter.New(ctx, driver)
	//	if err != nil {
	//		log.Fatalf("failed to initialize exporter: %v", err)
	//	}
	//
	//	// Tracer provider를 설정합니다.
	//	provider := sdktrace.NewTracerProvider(
	//		sdktrace.WithBatcher(exporter),
	//		// 트레이싱을 항상 캡처하도록 설정합니다.
	//		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	//	)
	//
	//	otel.SetTracerProvider(provider)
	//
	//	// Tracer를 가져옵니다.
	//	tracer := otel.Tracer("test-tracer")
	//
	//	// 샘플 트레이싱 작업을 수행합니다.
	//	ctx, span := tracer.Start(ctx, "test-span", trace.WithAttributes(semconv.ServiceNameKey.String("test-service")))
	//	time.Sleep(time.Second)
	//	span.End()
	//
	//	// 트레이스를 Collector로 보냅니다.
	//	err = provider.Shutdown(ctx)
	//	if err != nil {
	//		log.Fatalf("failed to stop provider: %v", err)
	//	}
}

func trace1() {
	//	// OpenTelemetry 컬렉터 엔드포인트 설정
	//	endpoint := "192.168.20.125:4317"
	//
	//	// OpenTelemetry OTLP(OpenTelemetry Protocol) exporter 생성
	//	exp, err := otlp.NewExporter(
	//		context.Background(),
	//		otlp.WithInsecure(), // 보안을 위해 TLS를 사용하지 않음
	//		otlp.WithAddress(endpoint),
	//	)
	//	if err != nil {
	//		log.Fatalf("failed to create exporter: %v", err)
	//	}
	//	defer exp.Stop()
	//
	//	// OpenTelemetry Tracer 생성
	//	tp := trace.NewTracerProvider(
	//		trace.WithBatcher(exp),
	//		trace.WithResource(resource.NewWithAttributes(resource.SchemaURL, "service.name", "example-service")),
	//	)
	//	defer tp.Shutdown()
	//
	//	otel.SetTracerProvider(tp)
	//
	//	tracer := otel.Tracer("example-tracer")
	//
	//	// 예제 span 시작
	//	_, span := tracer.Start(context.Background(), "example-span")
	//	defer span.End()
	//
	//	// 트레이싱 로직 수행
	//	time.Sleep(time.Second)
	//	log.Println("Tracing information sent to OpenTelemetry collector successfully!")
}
