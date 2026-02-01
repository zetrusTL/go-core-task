package customwg

import (
	"sync"
)

type SemaphoreWG struct {
	counter    int         
	mu         sync.Mutex   
	semaphore  chan struct{} 
	waitCalled bool         
}

func NewSemaphoreWG() *SemaphoreWG {
	return &SemaphoreWG{
		semaphore: make(chan struct{}, 1), 
	}
}

func (wg *SemaphoreWG) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	oldCounter := wg.counter
	wg.counter += delta

	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	if oldCounter > 0 && wg.counter == 0 && wg.waitCalled {
		select {
		case wg.semaphore <- struct{}{}:
		default:
		}
	}
}

func (wg *SemaphoreWG) Done() {
	wg.Add(-1)
}

func (wg *SemaphoreWG) Wait() {
	wg.mu.Lock()
	
	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}

	wg.waitCalled = true
	wg.mu.Unlock()

	<-wg.semaphore
	
	wg.mu.Lock()
	wg.waitCalled = false
	wg.mu.Unlock()
}

func (wg *SemaphoreWG) GetCounter() int {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	return wg.counter
}