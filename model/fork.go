package model

import (
	"time"
)

type Fork struct {
	NumberOfTimesUsed *int
	IsFree            *bool
	In                chan int
	Out               chan int
}

func (f *Fork) run() {
	for {
		*f.NumberOfTimesUsed++

		time.Sleep(1 * time.Second)
	}
}

func NewFork() Fork {
	isFree := true
	timesUsed := 0
	f := Fork{
		NumberOfTimesUsed: &timesUsed,
		IsFree:            &isFree,
		In:                make(chan int),
		Out:               make(chan int),
	}

	go f.run()

	return f
}
