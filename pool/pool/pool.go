package pool

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

/**
初始化10个长度的集合接受用户提交的
当active有空闲就消耗一个task
*/
type Pool struct {
	capacity int //队列长度
	active   chan struct{}
	tasks    chan Task
	wg       sync.WaitGroup //用于pool销毁时等待所有work退出
	quit     chan struct{}  //用于通知各个worker退出的信号struct
}

func New(capacity int) *Pool {
	if capacity < 0 {
		capacity = 10
	}
	if capacity > 100 {
		capacity = 100
	}
	p := &Pool{
		capacity: capacity,
		quit:     make(chan struct{}),
		tasks:    make(chan Task),
		active:   make(chan struct{}, capacity),
	}

	fmt.Println("pool start")
	go p.run()
	return p
}

var E = errors.New("quit")

func (p *Pool) Schedule(t Task) error {
	//向worker的队列里面放
	select {
	case <-p.quit:
		return E
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quit:
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
				fmt.Printf("worker[%03d]:recover panic[%s] and exist\n", idx, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]:start\n", idx)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d] exit.\n", idx)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", idx)
				t()
			}

		}
	}()
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Println("end")
}
