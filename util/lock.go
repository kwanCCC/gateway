package util

import "fmt"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mutex := &Mutex{make(chan struct{}, 1)}
	mutex.ch <- struct{}{}
	return mutex
}

func (mutex *Mutex) Lock() {
	<-mutex.ch
}

func (mutex *Mutex) Unlock() {
	select {
	case mutex.ch <- struct{}{}:
		fmt.Println("Unlock")
	default:
		panic("unlock of unlocked mutex")
	}
}

func (mutex *Mutex) TryLock() (status bool) {
	select {
	case <-mutex.ch:
		status = true
	default:
		status = false
	}
	return
}

func (mutex *Mutex) isLock() bool {
	return len(mutex.ch) == 0
}
