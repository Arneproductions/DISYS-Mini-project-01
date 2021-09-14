package model

import (
	"fmt"
	"sync"
)

type Fork struct {
	mu sync.Mutex
	id                int
	numberOfTimesUsed int
	isFree            bool
	In                chan int
	Out               chan int
}

func (f *Fork) communicate() {
	for {
		in := <-f.In

		switch in {
		case 0:
			f.Out <- f.id
		case 1:
			f.Out <- f.numberOfTimesUsed
		case 2:
			if f.isFree {
				f.Out <- 1
			} else {
				f.Out <- 0
			}
		}
	}
}

func (f *Fork) GetId() int {
	f.In <- 0
	return <- f.Out
}

func (f *Fork) GetTimesUsed() int {
	f.In <- 1
	return <- f.Out
}

func (f *Fork) GetFree() bool {
	f.In <- 2
	switch i := <-f.Out; i {
	default:
		fallthrough
	case 0:
		return false
	case 1:
		return true
	}
}

func (f *Fork) GetStatus() string {
	return fmt.Sprintf("Fork %d: Free: %t, Used: %d times", f.GetId(), f.GetFree(), f.GetTimesUsed())
}

func (f *Fork) Lock() {
	f.mu.Lock()
	f.isFree = false
}

func (f *Fork) Unlock() {
	f.numberOfTimesUsed++
	f.isFree = true
	f.mu.Unlock()
}

func NewFork(id int) *Fork {
	f := Fork{
		id:                id,
		numberOfTimesUsed: 0,
		isFree:            false,
		In:                make(chan int),
		Out:               make(chan int),
	}
	
	go f.communicate()

	return &f
}
