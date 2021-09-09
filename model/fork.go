package model

import (
	"time"
)

type Fork struct {
	NumberOfTimesUsed int
	IsFree            bool
	In                chan int
	Out               chan int
}

func (f *Fork) run() {
	for {
		f.NumberOfTimesUsed++

		time.Sleep(1 * time.Second)
	}
}

func NewFork() Fork {
	f := Fork{
		NumberOfTimesUsed: 0,
		IsFree:            true,
		In:                make(chan int),
		Out:               make(chan int),
	}

	go f.run()

	return f
}
