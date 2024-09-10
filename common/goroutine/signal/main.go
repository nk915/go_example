package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// exam01() // Not 시그널
	// exam02() // 시그널
	// exam03() // 시간 취소 컨텍스트
	exam04() //  컨텍스트 + 시그널
}

func exam04() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go printRoutineWithContext(ctx, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("exam04: receive sigint signal")

	cancel()
	wg.Wait()
	log.Println("exam04: ctx receive sigterm signal")

}

func exam03() {
	go printRoutine()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-ctx.Done()
	log.Println("receive sigint signal")
}

func exam02() {
	go printRoutine()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}

func exam01() {
	go printRoutine()
	time.Sleep(10 * time.Second)
}

func printRoutineWithContext(ctx context.Context, wg *sync.WaitGroup) {
	i := 0
	isComplete := false

	defer wg.Done()
	defer func() {
		log.Printf("routine: is complete print Routine : %v\n", isComplete)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("routine: ctx done printRoutine close")
			return
		default:
			i++
			log.Printf("routine: print Routine : %d\n", i)
			isComplete = false
			time.Sleep(1 * time.Second)
			isComplete = true
		}
	}
}

func printRoutine() {
	i := 0
	isComplete := false
	defer func() {
		log.Printf("is complete print Routine : %v\n", isComplete) // 안찍힘
	}()
	for {
		select {
		default:
			i++
			log.Printf("print Routine : %d\n", i)
			isComplete = false
			time.Sleep(1 * time.Second)
			isComplete = true
		}
	}
}
