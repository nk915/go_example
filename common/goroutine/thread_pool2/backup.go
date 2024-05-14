package backup

import (
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	ProcessMessage(msg string)
}

type A struct {
	Name string
}
type B struct{}
type C struct{}
type D struct{}

func (a A) ProcessMessage(msg string) {
	fmt.Printf("A(%s) received:%s\n", a.Name, msg)
	time.Sleep(3 * time.Second)
	fmt.Printf("A(%s) finsished:%s \n", a.Name, msg)
}

func (b B) ProcessMessage(msg string) { fmt.Println("B received:", msg) }
func (c C) ProcessMessage(msg string) { fmt.Println("C received:", msg) }
func (d D) ProcessMessage(msg string) { fmt.Println("D received:", msg) }

type Pool struct {
	workers map[string]chan string
	mutex   sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		workers: make(map[string]chan string),
	}
}

func (p *Pool) StartWorker(workerType string, worker Worker, count int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, exists := p.workers[workerType]; !exists {
		ch := make(chan string)
		p.workers[workerType] = ch
	}

	for i := 0; i < count; i++ {
		go func(c chan string) {
			for msg := range c {
				worker.ProcessMessage(msg)
			}
		}(p.workers[workerType])
	}
}

func (p *Pool) StopWorker(workerType string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if ch, exists := p.workers[workerType]; exists {
		close(ch)                     // Closes the channel to terminate the goroutine.
		delete(p.workers, workerType) // Removes the worker list of the corresponding type.
	}
}

func (p *Pool) SendMessage(workerType string, msg string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if ch, exists := p.workers[workerType]; exists {
		ch <- msg // Sends the message to the shared channel.
	}
}

func main() {
	pool := NewPool()

	pool.StartWorker("A", A{Name: "A-1"}, 1)
	pool.StartWorker("A", A{Name: "A-2"}, 1)
	pool.StartWorker("A", A{Name: "A-3"}, 1)
	pool.StartWorker("B", B{}, 3) // All B workers share the same channel now.

	time.Sleep(1 * time.Second) // Simulate time for some processing

	pool.SendMessage("A", "Message for 1A")
	pool.SendMessage("A", "Message for 2A")
	pool.SendMessage("A", "Message for 3A")
	pool.SendMessage("A", "Message for 4A")
	pool.SendMessage("A", "Message for 5A")
	pool.SendMessage("B", "Message for B")

	time.Sleep(1 * time.Second) // Allow some time for messages to be processed

	fmt.Println("Call Stop Worker A")
	pool.StopWorker("A")
	pool.StartWorker("C", C{}, 3)

	pool.SendMessage("C", "Message for C")

	time.Sleep(1 * time.Second) // Allow some time for messages to be processed

	// This simple test demonstrates starting, messaging, and stopping workers.
	// In a real application, you might need more sophisticated control and better message distribution.
}
