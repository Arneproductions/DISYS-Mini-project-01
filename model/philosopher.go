package model

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id         int
	right      *Fork
	left       *Fork
	eating     bool
	timesEaten int
	in         chan int
	out        chan int
}

func (p *Philosopher) communicate() {
	for {
		in := <-p.in

		switch in {
		case 0:
			p.out <- p.id
		case 1:
			p.out <- p.timesEaten
		case 2:
			if p.eating {
				p.out <- 1
			} else {
				p.out <- 0
			}
		}
	}
}

func (p *Philosopher) eat() {
	for {
		p.left.Lock()
		p.right.Lock()

		p.eating = true
		time.Sleep(5 * time.Millisecond)
		p.timesEaten++
		p.eating = false

		p.left.Unlock()
		p.right.Unlock()
	}
}

func (p *Philosopher) GetId() int {
	p.in <- 0
	return <- p.out
}

func (p *Philosopher) GetTimesEaten() int {
	p.in <- 1
	return <- p.out
}

func (p *Philosopher) GetEating() bool {
	p.in <- 2
	switch i := <-p.out; i {
	default:
		fallthrough
	case 0:
		return false
	case 1:
		return true
	}
}

func (p *Philosopher) GetStatus() string {
	return fmt.Sprintf("Philosopher %d: Eating: %t, Ate: %d times", p.GetId(), p.GetEating(), p.GetTimesEaten())
}

func NewPhilosopher(id int, right *Fork, left *Fork) Philosopher {
	p := Philosopher{
		id:         id,
		right:      right,
		left:       left,
		eating:     false,
		timesEaten: 0,
		in:         make(chan int),
		out:        make(chan int),
	}

	go p.eat()
	go p.communicate()

	return p
}
