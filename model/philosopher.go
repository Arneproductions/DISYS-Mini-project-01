package model

import (
	"time"
)

type philosopher struct {
	right fork 
	left fork
	eating bool
	timesEaten int 
	in chan int
	out chan int
}

func (p *philosopher) run() {
	for {
		p.timesEaten++

		time.Sleep(1 * time.Second)
	}
}

func NewPhilosopher (right fork, left fork) philosopher {
	p := philosopher{
		right: right,
		left: left,
		eating: false,
		timesEaten: 0,
		in: make(chan int),
		out: make(chan int),
	}

	go p.run()

	return p
}
