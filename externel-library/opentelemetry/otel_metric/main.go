package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var meter = otel.Meter("app_or_package_name")
var serviceName = semconv.ServiceNameKey.String("test-service")

// Initialize a gRPC connection to be used by both the tracer and meter
// providers.
func initConn() (*grpc.ClientConn, error) {
	// It connects the OpenTelemetry Collector through local gRPC connection.
	// You may replace `localhost:4317` with your endpoint.
	conn, err := grpc.NewClient("192.168.20.202:55682",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

// Initializes an OTLP exporter, and configures the corresponding meter provider.
func initMeterProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithGRPCConn(conn),
		// otlpmetricgrpc.WithCompressor(gzip.Name),
		otlpmetricgrpc.WithTemporalitySelector(preferDeltaTemporalitySelector),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
			//sdkmetric.WithInterval(5*time.Second),
			)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}

func preferDeltaTemporalitySelector(kind sdkmetric.InstrumentKind) metricdata.Temporality {
	switch kind {
	case sdkmetric.InstrumentKindCounter,
		sdkmetric.InstrumentKindObservableCounter,
		sdkmetric.InstrumentKindHistogram:
		return metricdata.DeltaTemporality
	default:
		return metricdata.CumulativeTemporality
	}
}

func main() {
	log.Printf("Waiting for connection...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := initConn()
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownMeterProvider, err := initMeterProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownMeterProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown MeterProvider: %s", err)
		} else {
			log.Printf("success to shutdown")
		}
	}()

	// Synchronous instruments.
	// upDownCounter(ctx)
	go counter(ctx)
	go upDownCounter(ctx)
	go histogram(ctx)
	//
	//	// Asynchronous instruments.
	//	go counterObserver(ctx)
	//	go upDownCounterObserver(ctx)
	//	go gaugeObserver(ctx)
	//
	//	// Advanced.
	//	go counterWithLabels(ctx)
	//	go counterObserverAdvanced(ctx)

	fmt.Println("reporting measurements to Uptrace... (press Ctrl+C to stop)")

	ch := make(chan os.Signal, 3)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
}

// counter demonstrates how to measure non-decreasing numbers, for example,
// number of requests or connections.
func counter(ctx context.Context) {
	counter, _ := meter.Int64Counter(
		"some.prefix.counter",
		metric.WithUnit("1"),
		metric.WithDescription("TODO"),
	)

	for {
		counter.Add(ctx, 1)
		time.Sleep(time.Second)
	}
}

// upDownCounter demonstrates how to measure numbers that can go up and down, for example,
// number of goroutines or customers.
func upDownCounter(ctx context.Context) {
	counter, _ := meter.Int64UpDownCounter(
		"some.prefix.up_down_counter",
		metric.WithUnit("1"),
		metric.WithDescription("TODO"),
	)

	for {
		if rand.Float64() >= 0.5 {
			counter.Add(ctx, +1)
		} else {
			counter.Add(ctx, -1)
		}

		time.Sleep(time.Second)
	}
}

// histogram demonstrates how to record a distribution of individual values, for example,
// request or query timings. With this instrument you get total number of records,
// avg/min/max values, and heatmaps/percentiles.
func histogram(ctx context.Context) {
	durRecorder, _ := meter.Int64Histogram(
		"some.prefix.histogram",
		metric.WithUnit("microseconds"),
		metric.WithDescription("TODO"),
	)

	for {
		dur := time.Duration(rand.NormFloat64()*5000000) * time.Microsecond
		durRecorder.Record(ctx, dur.Microseconds())

		time.Sleep(time.Millisecond)
	}
}

// counterObserver demonstrates how to measure monotonic (non-decreasing) numbers,
// for example, number of requests or connections.
func counterObserver(ctx context.Context) {
	counter, _ := meter.Int64ObservableCounter(
		"some.prefix.counter_observer",
		metric.WithUnit("1"),
		metric.WithDescription("TODO"),
	)

	var number int64
	if _, err := meter.RegisterCallback(
		// SDK periodically calls this function to collect data.
		func(ctx context.Context, o metric.Observer) error {
			number++
			o.ObserveInt64(counter, number)
			return nil
		},
	); err != nil {
		panic(err)
	}
}

// upDownCounterObserver demonstrates how to measure numbers that can go up and down,
// for example, number of goroutines or customers.
func upDownCounterObserver(ctx context.Context) {
	counter, err := meter.Int64ObservableUpDownCounter(
		"some.prefix.up_down_counter_async",
		metric.WithUnit("1"),
		metric.WithDescription("TODO"),
	)
	if err != nil {
		panic(err)
	}

	if _, err := meter.RegisterCallback(
		func(ctx context.Context, o metric.Observer) error {
			num := runtime.NumGoroutine()
			o.ObserveInt64(counter, int64(num))
			return nil
		},
		counter,
	); err != nil {
		panic(err)
	}
}

// gaugeObserver demonstrates how to measure non-additive numbers that can go up and down,
// for example, cache hit rate or memory utilization.
func gaugeObserver(ctx context.Context) {
	gauge, _ := meter.Float64ObservableGauge(
		"some.prefix.gauge_observer",
		metric.WithUnit("1"),
		metric.WithDescription("TODO"),
	)

	if _, err := meter.RegisterCallback(
		func(ctx context.Context, o metric.Observer) error {
			o.ObserveFloat64(gauge, rand.Float64())
			return nil
		},
		gauge,
	); err != nil {
		panic(err)
	}
}

// counterWithLabels demonstrates how to add different labels ("hits" and "misses")
// to measurements. Using this simple trick, you can get number of hits, misses,
// sum = hits + misses, and hit_rate = hits / (hits + misses).
func counterWithLabels(ctx context.Context) {
	counter, _ := meter.Int64Counter(
		"some.prefix.cache",
		metric.WithDescription("Cache hits and misses"),
	)
	for {
		if rand.Float64() < 0.3 {
			// increment hits
			counter.Add(ctx, 1, metric.WithAttributes(attribute.String("type", "hits")))
		} else {
			// increments misses
			counter.Add(ctx, 1, metric.WithAttributes(attribute.String("type", "misses")))
		}

		time.Sleep(time.Millisecond)
	}
}

// counterObserverAdvanced demonstrates how to measure monotonic (non-decreasing) numbers,
// for example, number of requests or connections.
func counterObserverAdvanced(ctx context.Context) {
	// stats is our data source updated by some library.
	var stats struct {
		Hits   int64 // atomic
		Misses int64 // atomic
	}

	hitsCounter, _ := meter.Int64ObservableCounter("some.prefix.cache_hits")
	missesCounter, _ := meter.Int64ObservableCounter("some.prefix.cache_misses")

	if _, err := meter.RegisterCallback(
		// SDK periodically calls this function to collect data.
		func(ctx context.Context, o metric.Observer) error {
			o.ObserveInt64(hitsCounter, atomic.LoadInt64(&stats.Hits))
			o.ObserveInt64(missesCounter, atomic.LoadInt64(&stats.Misses))
			return nil
		},
		hitsCounter,
		missesCounter,
	); err != nil {
		panic(err)
	}

	for {
		if rand.Float64() < 0.3 {
			atomic.AddInt64(&stats.Misses, 1)
		} else {
			atomic.AddInt64(&stats.Hits, 1)
		}

		time.Sleep(time.Millisecond)
	}
}
