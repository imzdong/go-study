package main

import (
	"fmt"
	"pool/pool"
	"time"
)

func main()  {
	pool1()
}

func pool1() {
	
	
	p := pool.New(5)
	fmt.Println("6666")
	time.Sleep(5*time.Second)
	
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
			fmt.Printf("i'm task%d\n", i)
		})
		if err != nil {
			fmt.Println(err)
		}
	}

	p.Free()
	fmt.Println("main end")
}
