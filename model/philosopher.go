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
	In         chan int
	Out        chan int
}

func (p *Philosopher) communicate() {
	for {
		in := <-p.In

		switch in {
		case 0:
			p.Out <- p.id
		case 1:
			p.Out <- p.timesEaten
		case 2:
			if p.eating {
				p.Out <- 1
			} else {
				p.Out <- 0
			}
		}
	}
}

func (p *Philosopher) eat() {
	for {
		p.left.Lock()
		p.right.Lock()

		p.eating = true
		time.Sleep(1 * time.Second)
		p.timesEaten++
		p.eating = false

		p.left.Unlock()
		p.right.Unlock()
	}
}

func (p *Philosopher) GetId() int {
	p.In <- 0
	return <- p.Out
}

func (p *Philosopher) GetTimesEaten() int {
	p.In <- 1
	return <- p.Out
}

func (p *Philosopher) GetEating() bool {
	p.In <- 2
	switch i := <-p.Out; i {
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
		In:         make(chan int),
		Out:        make(chan int),
	}

	go p.eat()
	go p.communicate()

	return p
}
