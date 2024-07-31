package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	ProcessMessage(msg string)
}

type A struct{ Name string }
type B struct{}

func (a A) ProcessMessage(msg string) {
	fmt.Printf("A(%s) received:%s\n", a.Name, msg)
	time.Sleep(3 * time.Second)
	fmt.Printf("A(%s) finsished:%s \n", a.Name, msg)
}
func (b B) ProcessMessage(msg string) { fmt.Println("B received:", msg) }

type Pool struct {
	workers      map[string]chan string
	stopChannels map[string]chan bool
	waitGroups   map[string]*sync.WaitGroup // 각 워커 타입별로 WaitGroup을 저장합니다.
	mutex        sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		workers:      make(map[string]chan string),
		stopChannels: make(map[string]chan bool),
		waitGroups:   make(map[string]*sync.WaitGroup),
	}
}

func (p *Pool) StartWorker(workerType string, worker Worker, count int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, exists := p.workers[workerType]; !exists {
		workerCh := make(chan string)
		stopCh := make(chan bool)
		wg := &sync.WaitGroup{} // WaitGroup 인스턴스를 생성합니다.
		p.workers[workerType] = workerCh
		p.stopChannels[workerType] = stopCh
		p.waitGroups[workerType] = wg
		for i := 0; i < count; i++ {
			wg.Add(1) // WaitGroup에 고루틴 추가를 알립니다.
			go func(c chan string, stop chan bool, wg *sync.WaitGroup, id int) {
				defer wg.Done() // 고루틴이 종료될 때 Done을 호출합니다.
				for {
					select {
					case msg := <-c:
						worker.ProcessMessage(fmt.Sprintln(msg, id))
					case <-stop:
						fmt.Println("stop: ", id)
						return
					}
				}
			}(workerCh, stopCh, wg, i)
		}
	}
}

func (p *Pool) Wait(workerType string) {
	p.waitGroups[workerType].Wait() // 모든 고루틴이 종료될 때까지 대기합니다.
}

func (p *Pool) StopWorker(workerType string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if stopCh, exists := p.stopChannels[workerType]; exists {
		close(stopCh)                   // 종료 신호를 보냅니다.
		p.waitGroups[workerType].Wait() // 모든 고루틴이 종료될 때까지 대기합니다.
		delete(p.stopChannels, workerType)
		if ch, exists := p.workers[workerType]; exists {
			close(ch)
			delete(p.workers, workerType)
			delete(p.waitGroups, workerType) // WaitGroup을 삭제합니다.
		}
	}
}

func (p *Pool) SendMessage(workerType string, msg string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if ch, exists := p.workers[workerType]; exists {
		ch <- msg
	}
}

func main() {
	pool := NewPool()

	pool.StartWorker("A", A{Name: "1"}, 3)
	pool.StartWorker("A", A{Name: "1"}, 2)
	//pool.StartWorker("B", B{}, 3)

	time.Sleep(1 * time.Second)

	pool.SendMessage("A", "Message for 1A")
	pool.SendMessage("A", "Message for 2A")
	pool.SendMessage("A", "Message for 3A")
	pool.SendMessage("A", "Message for 4A")
	pool.SendMessage("A", "Message for 5A")
	//pool.SendMessage("B", "Message for B")

	time.Sleep(time.Second / 2)

	fmt.Println("A workers stopping")
	pool.StopWorker("A") // A 타입 워커를 종료합니다. 모든 메시지 처리가 완료될 때까지 대기합니다.
	fmt.Println("A workers stopped")

	//	pool.StartWorker("C", B{}, 3) // B 타입 워커를 C로 재시작합니다. (예제에서는 같은 B 인스턴스 사용)
	//
	//	pool.SendMessage("C", "Message for C")
	//
	//	fmt.Println("C workers stopping")
	time.Sleep(15 * time.Second)
	//
	//	//pool.StopWorker("C")
	//	fmt.Println("C workers stopped")
}
