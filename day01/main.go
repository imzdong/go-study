package main

import "fmt"

func main() {
	day01DataVar()
}

// 常量
const (
	c1 = 6
	c2 = 7
) 

func day01DataVar() {
	fmt.Println(c1)
	var a int = 10
	fmt.Println("单一声明")
	fmt.Println(a)
	var (
		b int = 6
		c string = "imzdong"
	)
	fmt.Println("区块声明")
	fmt.Println(b)
	fmt.Println(c)
	
	var d ,e string = "7", "sex"
	fmt.Println("一行声明多个")
	fmt.Println(d)
	fmt.Println(e)
	
	var(
		f,g int = 8, 9
	)
	fmt.Println("区块声明单一声明多个")
	fmt.Println(f)
	fmt.Println(g)
	
	var h = 7
	fmt.Printf("单一声明无显示类型：%d\n", h)
	// 局部变量建议
	i := 8
	fmt.Printf("短声明：%d\n", i)
}