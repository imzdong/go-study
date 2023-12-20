package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{},2)
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
