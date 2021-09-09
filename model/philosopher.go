package model

import (
	"fmt"
	"time"
)

type Philosopher struct {
	right      Fork
	left       Fork
	eating     bool
	timesEaten int
	in         chan int
	out        chan int
}

func (p *Philosopher) run() {
	for {
		p.timesEaten++

		time.Sleep(1 * time.Second)
	}
}

func (p *Philosopher) GetStatus() string {
	return fmt.Sprintf("Is Eating: %t, Ate: %d", p.eating, p.timesEaten)
}

func NewPhilosopher(right Fork, left Fork) Philosopher {
	p := Philosopher{
		right:      right,
		left:       left,
		eating:     false,
		timesEaten: 0,
		in:         make(chan int),
		out:        make(chan int),
	}

	go p.run()

	return p
}
