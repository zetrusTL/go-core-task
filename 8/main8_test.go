package customwg

import (
	"sync"
	"testing"
	"time"
)

func TestNewSemaphoreWG(t *testing.T) {
	wg := NewSemaphoreWG()
	
	if wg == nil {
		t.Fatal("NewSemaphoreWG вернул nil")
	}
	
	if wg.GetCounter() != 0 {
		t.Errorf("Ожидался счетчик 0, получил %d", wg.GetCounter())
	}
}

func TestAdd(t *testing.T) {
	wg := NewSemaphoreWG()
	
	wg.Add(3)
	if wg.GetCounter() != 3 {
		t.Errorf("Ожидался счетчик 3, получил %d", wg.GetCounter())
	}
	
	wg.Add(-2)
	if wg.GetCounter() != 1 {
		t.Errorf("Ожидался счетчик 1, получил %d", wg.GetCounter())
	}
}

func TestAddPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Ожидалась паника при отрицательном счетчике")
		}
	}()
	
	wg := NewSemaphoreWG()
	wg.Add(-1) 
}

func TestDone(t *testing.T) {
	wg := NewSemaphoreWG()
	wg.Add(3)
	
	wg.Done()
	if wg.GetCounter() != 2 {
		t.Errorf("Ожидался счетчик 2, получил %d", wg.GetCounter())
	}
	
	wg.Done()
	wg.Done()
	if wg.GetCounter() != 0 {
		t.Errorf("Ожидался счетчик 0, получил %d", wg.GetCounter())
	}
}

func TestWait(t *testing.T) {
	wg := NewSemaphoreWG()
	
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	
	if elapsed > 50*time.Millisecond {
		t.Error("Wait не должен блокироваться при счетчике 0")
	}
}

func TestWaitWithGoroutines(t *testing.T) {
	wg := NewSemaphoreWG()
	const numGoroutines = 5
	
	completed := make([]bool, numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			completed[idx] = true
		}(i)
	}
	
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	
	if elapsed < 5*time.Millisecond {
		t.Error("Wait завершился слишком быстро, не дождавшись горутин")
	}
	
	for i, done := range completed {
		if !done {
			t.Errorf("Горутина %d не завершилась", i)
		}
	}
	
	if wg.GetCounter() != 0 {
		t.Errorf("Ожидался счетчик 0 после Wait, получил %d", wg.GetCounter())
	}
}

func TestConcurrentAccess(t *testing.T) {
	wg := NewSemaphoreWG()
	const numWorkers = 100
	
	var wgMain sync.WaitGroup
	wgMain.Add(numWorkers * 2)
	
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wgMain.Done()
			wg.Add(1)
			time.Sleep(time.Microsecond)
		}()
	}
	
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wgMain.Done()
			time.Sleep(time.Microsecond)
			wg.Done()
		}()
	}
	
	wgMain.Wait()
	
	if wg.GetCounter() != 0 {
		t.Errorf("Ожидался счетчик 0 после конкурентного доступа, получил %d", wg.GetCounter())
	}
}

func TestMultipleWait(t *testing.T) {
	wg := NewSemaphoreWG()
	const numGoroutines = 3
	
	waitCompleted := make(chan bool, 5)
	
	for i := 0; i < 5; i++ {
		go func() {
			wg.Wait()
			waitCompleted <- true
		}()
	}
	
	time.Sleep(10 * time.Millisecond)
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(20 * time.Millisecond)
		}()
	}
	
	timeout := time.After(100 * time.Millisecond)
	completedCount := 0
	
	for i := 0; i < 5; i++ {
		select {
		case <-waitCompleted:
			completedCount++
		case <-timeout:
			t.Fatalf("Таймаут ожидания Wait, завершено %d из 5", completedCount)
		}
	}
	
	if completedCount != 5 {
		t.Errorf("Не все Wait завершились: %d из 5", completedCount)
	}
}

func TestReuse(t *testing.T) {
	wg := NewSemaphoreWG()
	
	wg.Add(2)
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Done()
	}()
	
	go func() {
		time.Sleep(20 * time.Millisecond)
		wg.Done()
	}()
	
	wg.Wait()
	
	wg.Add(1)
	
	completed := false
	go func() {
		time.Sleep(10 * time.Millisecond)
		completed = true
		wg.Done()
	}()
	
	wg.Wait()
	
	if !completed {
		t.Error("WaitGroup не сработал при повторном использовании")
	}
}