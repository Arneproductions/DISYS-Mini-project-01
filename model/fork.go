package model

import (
	"time"
)

type fork struct {
	NumberOfTimesUsed int
	IsFree bool
	In chan int
	Out chan int
}

func (f *fork) run() {
	for {
		f.NumberOfTimesUsed++

		time.Sleep(1 * time.Second)
	}
}

func NewFork() fork {
	f := fork {
		NumberOfTimesUsed: 0,
		IsFree: true,
		In: make(chan int),
		Out: make(chan int),
	}

	go f.run()

	return f
}
