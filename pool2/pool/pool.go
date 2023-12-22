package pool

import (
	"errors"
	"fmt"
	"sync"
)

type Pool struct {
	capcity int
	tasks   chan Task
	active  chan struct{}
	quit    chan struct{}
	wg      sync.WaitGroup
}

type Task func()

func New(capcity int) *Pool {
	if capcity < 0 {
		capcity = 10
	}
	if capcity > 100 {
		capcity = 100
	}
	p := &Pool{
		capcity: capcity,
		tasks:   make(chan Task),
		active:  make(chan struct{}, capcity),
		quit:    make(chan struct{}),
	}
	fmt.Println("init pool")
	go p.run()
	fmt.Println("init pool end")
	return p
}

func (p *Pool) run() {
	idx := 0
	fmt.Println("pool run")
	for {
		select {
		case <-p.quit:
			fmt.Println("quit")
			<-p.active
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(idx int) {

	p.wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker:[%03d] error,%s\n", idx, err)
				<-p.active
			}
			p.wg.Done()
			fmt.Printf("worker:[%03d] done\n", idx)
		}()

		fmt.Printf("worker:[%03d] start\n", idx)
		select {
		case <-p.quit:
			fmt.Printf("worker:[%03d] start exit\n", idx)
			<-p.active
			return
		case t := <-p.tasks:
			fmt.Printf("worker:[%03d] start working\n", idx)
			t()
			fmt.Printf("worker:[%03d] start work end\n", idx)
		}
	}()

}

var exitError error = errors.New("pool exit")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return exitError
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit) // make sure all worker and p.run exit and schedule return error
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}
