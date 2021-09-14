package model

import (
	"sync"
)

type Fork struct {
	mu sync.Mutex
	Id                int
	NumberOfTimesUsed *int
	IsFree            *bool
	In                chan int
	Out               chan int
}


func NewFork(id int) Fork {
	isFree := true
	timesUsed := 0
	f := Fork{
		Id:                id,
		NumberOfTimesUsed: &timesUsed,
		IsFree:            &isFree,
		In:                make(chan int),
		Out:               make(chan int),
	}

	return f
}
