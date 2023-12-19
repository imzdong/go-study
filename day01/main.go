package main

import (
	"fmt"
	"github.com/imzdong/day01/bdata"
	"sync"
	"time"
)

func produce(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("p c<-")
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}
func consume(c <-chan int) {
	for v := range c {
		fmt.Println("v <-c")
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("produce start")
		produce(ch)
		fmt.Println("ppp")
		wg.Done()
		fmt.Println("p down")
	}()
	go func() {
		fmt.Println("consum start")
		consume(ch)
		fmt.Println("ccc")
		wg.Done()
		fmt.Println("c down")
	}()
	fmt.Println("main wait")
	wg.Wait()
	fmt.Println("main end")
}

func ch() {
	ch1 := make(chan int)
	go func() {
		fmt.Println("start sub")
		ch1 <- 13
		fmt.Println("end sub")
	}()
	fmt.Println("main start revice")
	n := <-ch1
	fmt.Println(n)
}

func study() {
	var bd bdata.Bdata
	fmt.Println(bd.Name)
	bd.Name = "winter"
	bd.Age = 18
	fmt.Printf("bd type:%T\n", bd)
	fmt.Println(bd)
	fmt.Println(&bd)
	bd1 := bdata.Bdata{Name: "w", Age: 18}
	fmt.Printf("bd1 type:%T\n", bd1)
	//fmt.Printf("bd==nil:%t\n", (bd==nil))

	bd2 := new(bdata.Bdata)
	fmt.Printf("bd2 type:%T\n", bd2)
	var bdp *bdata.Bdata
	fmt.Printf("bd2==bdp:%t\n", (bd2 == nil))

	fmt.Println(bd2)
	bd2.Name = "6x"
	fmt.Println(bd2)

	fmt.Printf("bdp type:%T\n", bdp)
	fmt.Println(bdp)
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal 0xc0000005 code=0x0 addr=0x0 pc=0xd0cf3d]
	//fmt.Println(bdp.Name)
	fmt.Printf("bdp==nil type:%t\n", (bdp == nil))
	//bdata.Day01DataVar()
	//bdata.Day01DataNum()
	//day01DataStr()
	//day01Array()
	//day02Map()
	//bd.Day02Struct()
}
