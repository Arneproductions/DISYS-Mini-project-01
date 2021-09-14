package model

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id         int
	right      *Fork
	left       *Fork
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
		time.Sleep(5 * time.Millisecond)
	}
}

func (p *Philosopher) eat() {
	for {
		p.left.mu.Lock()
		p.right.mu.Lock()

		*p.timesEaten++
		time.Sleep(1 * time.Second)

		// TODO: fork times used
		p.left.mu.Unlock()
		p.right.mu.Unlock()
	}
}

func (p *Philosopher) GetStatus() string {
	return fmt.Sprintf("phil %d: Is Eating: %t, Ate: %d times", p.id, *p.eating, *p.timesEaten)
}

func NewPhilosopher(id int, right *Fork, left *Fork) Philosopher {
	eating := false
	timesEaten := 0
	p := Philosopher{
		id:         id,
		right:      right,
		left:       left,
		eating:     &eating,
		timesEaten: &timesEaten,
		in:         make(chan int),
		out:        make(chan int),
	}

	go p.eat()

	return p
}
