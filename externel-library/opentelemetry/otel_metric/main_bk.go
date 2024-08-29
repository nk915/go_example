//package main
//
//import (
//	"context"
//	"fmt"
//	"math/rand"
//	"os"
//	"os/signal"
//	"runtime"
//	"sync/atomic"
//	"syscall"
//	"time"
//
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/attribute"
//	"go.opentelemetry.io/otel/metric"
//
//	"github.com/uptrace/uptrace-go/uptrace"
//)
//
//var meter = otel.Meter("app_or_package_name")
//
//func main() {
//	ctx := context.Background()
//
//	// Configure OpenTelemetry with sensible defaults.
//	uptrace.ConfigureOpentelemetry(
//		// copy your project DSN here or use UPTRACE_DSN env var
//		// uptrace.WithDSN("https://<key>@api.uptrace.dev/<project_id>"),
//
//		uptrace.WithServiceName("myservice"),
//		uptrace.WithServiceVersion("1.0.0"),
//	)
//	// Send buffered spans and free resources.
//	defer uptrace.Shutdown(ctx)
//
//	// Synchronous instruments.
//	go counter(ctx)
//	go upDownCounter(ctx)
//	go histogram(ctx)
//
//	// Asynchronous instruments.
//	go counterObserver(ctx)
//	go upDownCounterObserver(ctx)
//	go gaugeObserver(ctx)
//
//	// Advanced.
//	go counterWithLabels(ctx)
//	go counterObserverAdvanced(ctx)
//
//	fmt.Println("reporting measurements to Uptrace... (press Ctrl+C to stop)")
//
//	ch := make(chan os.Signal, 3)
//	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
//	<-ch
//}
//
//// counter demonstrates how to measure non-decreasing numbers, for example,
//// number of requests or connections.
//func counter(ctx context.Context) {
//	counter, _ := meter.Int64Counter(
//		"some.prefix.counter",
//		metric.WithUnit("1"),
//		metric.WithDescription("TODO"),
//	)
//
//	for {
//		counter.Add(ctx, 1)
//		time.Sleep(time.Millisecond)
//	}
//}
//
//// upDownCounter demonstrates how to measure numbers that can go up and down, for example,
//// number of goroutines or customers.
//func upDownCounter(ctx context.Context) {
//	counter, _ := meter.Int64UpDownCounter(
//		"some.prefix.up_down_counter",
//		metric.WithUnit("1"),
//		metric.WithDescription("TODO"),
//	)
//
//	for {
//		if rand.Float64() >= 0.5 {
//			counter.Add(ctx, +1)
//		} else {
//			counter.Add(ctx, -1)
//		}
//
//		time.Sleep(time.Second)
//	}
//}
//
//// histogram demonstrates how to record a distribution of individual values, for example,
//// request or query timings. With this instrument you get total number of records,
//// avg/min/max values, and heatmaps/percentiles.
//func histogram(ctx context.Context) {
//	durRecorder, _ := meter.Int64Histogram(
//		"some.prefix.histogram",
//		metric.WithUnit("microseconds"),
//		metric.WithDescription("TODO"),
//	)
//
//	for {
//		dur := time.Duration(rand.NormFloat64()*5000000) * time.Microsecond
//		durRecorder.Record(ctx, dur.Microseconds())
//
//		time.Sleep(time.Millisecond)
//	}
//}
//
//// counterObserver demonstrates how to measure monotonic (non-decreasing) numbers,
//// for example, number of requests or connections.
//func counterObserver(ctx context.Context) {
//	counter, _ := meter.Int64ObservableCounter(
//		"some.prefix.counter_observer",
//		metric.WithUnit("1"),
//		metric.WithDescription("TODO"),
//	)
//
//	var number int64
//	if _, err := meter.RegisterCallback(
//		// SDK periodically calls this function to collect data.
//		func(ctx context.Context, o metric.Observer) error {
//			number++
//			o.ObserveInt64(counter, number)
//			return nil
//		},
//	); err != nil {
//		panic(err)
//	}
//}
//
//// upDownCounterObserver demonstrates how to measure numbers that can go up and down,
//// for example, number of goroutines or customers.
//func upDownCounterObserver(ctx context.Context) {
//	counter, err := meter.Int64ObservableUpDownCounter(
//		"some.prefix.up_down_counter_async",
//		metric.WithUnit("1"),
//		metric.WithDescription("TODO"),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	if _, err := meter.RegisterCallback(
//		func(ctx context.Context, o metric.Observer) error {
//			num := runtime.NumGoroutine()
//			o.ObserveInt64(counter, int64(num))
//			return nil
//		},
//		counter,
//	); err != nil {
//		panic(err)
//	}
//}
//
//// gaugeObserver demonstrates how to measure non-additive numbers that can go up and down,
//// for example, cache hit rate or memory utilization.
//func gaugeObserver(ctx context.Context) {
//	gauge, _ := meter.Float64ObservableGauge(
//		"some.prefix.gauge_observer",
//		metric.WithUnit("1"),
//		metric.WithDescription("TODO"),
//	)
//
//	if _, err := meter.RegisterCallback(
//		func(ctx context.Context, o metric.Observer) error {
//			o.ObserveFloat64(gauge, rand.Float64())
//			return nil
//		},
//		gauge,
//	); err != nil {
//		panic(err)
//	}
//}
//
//// counterWithLabels demonstrates how to add different labels ("hits" and "misses")
//// to measurements. Using this simple trick, you can get number of hits, misses,
//// sum = hits + misses, and hit_rate = hits / (hits + misses).
//func counterWithLabels(ctx context.Context) {
//	counter, _ := meter.Int64Counter(
//		"some.prefix.cache",
//		metric.WithDescription("Cache hits and misses"),
//	)
//	for {
//		if rand.Float64() < 0.3 {
//			// increment hits
//			counter.Add(ctx, 1, metric.WithAttributes(attribute.String("type", "hits")))
//		} else {
//			// increments misses
//			counter.Add(ctx, 1, metric.WithAttributes(attribute.String("type", "misses")))
//		}
//
//		time.Sleep(time.Millisecond)
//	}
//}
//
//// counterObserverAdvanced demonstrates how to measure monotonic (non-decreasing) numbers,
//// for example, number of requests or connections.
//func counterObserverAdvanced(ctx context.Context) {
//	// stats is our data source updated by some library.
//	var stats struct {
//		Hits   int64 // atomic
//		Misses int64 // atomic
//	}
//
//	hitsCounter, _ := meter.Int64ObservableCounter("some.prefix.cache_hits")
//	missesCounter, _ := meter.Int64ObservableCounter("some.prefix.cache_misses")
//
//	if _, err := meter.RegisterCallback(
//		// SDK periodically calls this function to collect data.
//		func(ctx context.Context, o metric.Observer) error {
//			o.ObserveInt64(hitsCounter, atomic.LoadInt64(&stats.Hits))
//			o.ObserveInt64(missesCounter, atomic.LoadInt64(&stats.Misses))
//			return nil
//		},
//		hitsCounter,
//		missesCounter,
//	); err != nil {
//		panic(err)
//	}
//
//	for {
//		if rand.Float64() < 0.3 {
//			atomic.AddInt64(&stats.Misses, 1)
//		} else {
//			atomic.AddInt64(&stats.Hits, 1)
//		}
//
//		time.Sleep(time.Millisecond)
//	}
//}