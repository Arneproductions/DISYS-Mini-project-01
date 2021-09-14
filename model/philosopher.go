package model

import (
	"fmt"
	"time"
)

type Philosopher struct {
	right      Fork
	left       Fork
	eating     *bool
	timesEaten *int
	in         chan int
	out        chan int
}

func (p *Philosopher) run() {
	for {
		p.right.In <- 1

		*p.timesEaten = *p.timesEaten + 1
		*p.eating = !*p.eating
		time.Sleep(1 * time.Second)
	}
}

func (p *Philosopher) GetStatus() string {
	return fmt.Sprintf("Is Eating: %t, Ate: %d times", *p.eating, *p.timesEaten)
}

func NewPhilosopher(right Fork, left Fork) Philosopher {
	eating := false
	timesEaten := 0
	p := Philosopher{
		right:      right,
		left:       left,
		eating:     &eating,
		timesEaten: &timesEaten,
		in:         make(chan int),
		out:        make(chan int),
	}

	go p.run()

	return p
}
