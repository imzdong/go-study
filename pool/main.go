package main

import (
	"fmt"
	"imzdong.com/pool/pool"
	"time"
)

func main() {
	p := pool.New(5)

	time.Sleep(5 * time.Second)
	fmt.Printf("main sleep\n")

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			fmt.Printf("i'm task%d\n", i)
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			fmt.Println("main error")
		}

		fmt.Printf("main:%d\n", i)
	}

	p.Free()
}

func chtest() {
	done := make(chan struct{}, 2)
	fmt.Println("start go")
	go func() { worker(done) }()
	fmt.Println("end go, start done")
	<-done
	fmt.Println("end")
	<-done
	fmt.Println("end2")
	time.Sleep(4 * time.Second)
	fmt.Println("end4")
}

func worker(done chan struct{}) {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done working.")
	done <- struct{}{} // 发送空结构体表示工作完成
	time.Sleep(2 * time.Second)
	fmt.Println("Done working. 2")
	done <- struct{}{}
}
