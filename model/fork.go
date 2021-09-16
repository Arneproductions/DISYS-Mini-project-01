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
	in                chan int
	out               chan int
}

func (f *Fork) communicate() {
	for {
		in := <-f.in

		switch in {
		case 0:
			f.out <- f.id
		case 1:
			f.out <- f.numberOfTimesUsed
		case 2:
			if f.isFree {
				f.out <- 1
			} else {
				f.out <- 0
			}
		}
	}
}

func (f *Fork) GetId() int {
	f.in <- 0
	return <- f.out
}

func (f *Fork) GetTimesUsed() int {
	f.in <- 1
	return <- f.out
}

func (f *Fork) GetFree() bool {
	f.in <- 2
	switch i := <-f.out; i {
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
		in:                make(chan int),
		out:               make(chan int),
	}
	
	go f.communicate()

	return &f
}
