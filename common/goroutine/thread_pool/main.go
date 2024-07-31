package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Worker interface {
	Work(string)
	Recv(string, string)
	Name() string
}

// A, B, C 타입 정의
type A struct{}

func NewA() Worker               { return &A{} }
func (a *A) Work(id string)      { fmt.Printf("A 작업 수행: %s\n", id) }
func (a *A) Recv(id, msg string) { fmt.Printf("A (%s) : %s \n", id, msg) }
func (a *A) Name() string        { return "A" }

type B struct{}

func NewB() Worker               { return &B{} }
func (b *B) Work(id string)      { fmt.Printf("B 작업 수행: %s\n", id) }
func (b *B) Recv(id, msg string) { fmt.Printf("B (%s) : %s \n", id, msg) }
func (b *B) Name() string        { return "B" }

type C struct{}

func NewC() Worker               { return &C{} }
func (c *C) Work(id string)      { fmt.Printf("C 작업 수행: %s\n", id) }
func (c *C) Recv(id, msg string) { fmt.Printf("C (%s) : %s \n", id, msg) }
func (c *C) Name() string        { return "C" }

// ThreadPool 관리 구조체
type ThreadPool struct {
	mu        sync.Mutex
	active    map[string]int
	msgChan   map[string]chan string
	stopChans map[string][]chan struct{}
}

func NewThreadPool() *ThreadPool {
	return &ThreadPool{
		active:    make(map[string]int),
		msgChan:   make(map[string]chan string),
		stopChans: make(map[string][]chan struct{}),
	}
}

// Run은 지정된 Worker 타입에 대해 3개의 고루틴을 실행합니다.
func (tp *ThreadPool) Run(w Worker) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if _, ok := tp.active[w.Name()]; !ok {
		tp.active[w.Name()] = 0
	}
	if tp.active[w.Name()] >= 3 {
		fmt.Println(w.Name(), "작업은 이미 최대 수치로 실행 중입니다.")
		return
	}

	tp.msgChan[w.Name()] = make(chan string)

	for i := tp.active[w.Name()]; i < 3; i++ {
		stop := make(chan struct{})
		if tp.stopChans[w.Name()] == nil {
			tp.stopChans[w.Name()] = make([]chan struct{}, 0)
		}
		tp.stopChans[w.Name()] = append(tp.stopChans[w.Name()], stop)
		tp.active[w.Name()]++
		go tp.worker(w, tp.msgChan[w.Name()], stop)
	}

}

func (tp *ThreadPool) worker(w Worker, msg chan string, stop chan struct{}) {
	id := uuid.New().String()

	defer func() {
		tp.mu.Lock()
		tp.active[w.Name()]--
		if tp.active[w.Name()] == 0 {
			delete(tp.active, w.Name())
			delete(tp.stopChans, w.Name())
		}
		tp.mu.Unlock()
	}()

	for {
		select {
		case <-stop:
			fmt.Printf("%s 작업 중지: %s\n", w.Name(), id)
			return

		case m := <-msg:
			w.Recv(id, m)
			time.Sleep(1 * time.Second)
		default:
			w.Work(id)
			time.Sleep(1 * time.Second)
		}
	}
}

// Stop은 지정된 Worker 타입의 모든 고루틴을 중지합니다.
func (tp *ThreadPool) Stop(name string) {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	if chans, ok := tp.stopChans[name]; ok {
		for _, ch := range chans {
			close(ch)
		}
		fmt.Println(name, "작업을 중지합니다.")
	}
}

func (tp *ThreadPool) SendMessage(workerType string, msg string) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if ch, exists := tp.msgChan[workerType]; exists {
		ch <- msg
	}
}

func main() {
	pool := NewThreadPool()

	pool.Run(NewA())
	pool.Run(NewB())
	time.Sleep(2 * time.Second) // A와 B가 작업을 수행하는 동안 기다림

	pool.SendMessage("A", "1")
	pool.SendMessage("A", "11")
	pool.SendMessage("A", "111")
	pool.SendMessage("B", "2")
	pool.SendMessage("B", "22")

	//pool.Stop("A")

	time.Sleep(1 * time.Second)

	//pool.Run(NewB()) // B는 이미 실행 중이므로 추가 실행되지 않음

	time.Sleep(5 * time.Second)

	//	time.Sleep(5 * time.Second) // C가 작업을 수행하는 동안 기다림
}
