package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//day01DataVar()
	//day01DataNum()
	day01DataStr()
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
		b int    = 6
		c string = "imzdong"
	)
	fmt.Println("区块声明")
	fmt.Println(b)
	fmt.Println(c)

	var d, e string = "7", "sex"
	fmt.Println("一行声明多个")
	fmt.Println(d)
	fmt.Println(e)

	var (
		f, g int = 8, 9
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

func day01DataNum() {
	a8 := int8(1)
	fmt.Printf("int8:%b\n", a8)

	a16 := int16(16)
	fmt.Printf("int16:%d\n", a16)

	a32 := int32(32)
	fmt.Printf("int32:%d\n", a32)

	a64 := int64(64)
	fmt.Printf("int64:%d\n", a64)

	a := 6
	fmt.Printf("int:%d\n", a)

	f := float32(32)
	fmt.Printf("float32:%f\n", f)

	f64 := float64(64)
	fmt.Printf("float64:%f\n", f64)

	type myInt int
	m := myInt(6)
	fmt.Printf("myInt:%d\n", m)
	i := 6
	//(mismatched types myInt and int)
	//fmt.Printf("myInt==i:%d\n", (m==i))

	type myInt2 = int
	m2 := myInt2(6)
	m3 := myInt2(7)
	fmt.Printf("myInt2==i:%t\n", (m2 == i))
	fmt.Printf("myInt2==i:%t\n", (m3 == i))
}

func day01DataStr() {
	s := "6中国uuff,,..，，；‘"
	fmt.Printf("str byte length:%d\n", len(s))
	fmt.Printf("str char length:%d\n", utf8.RuneCountInString(s))
	c := rune('中')
	fmt.Printf("char :%d\n", c)
	fmt.Printf("char char :%T\n", c)

	sr := `所见即所得
			还支持换行
			厉害了呀`
	fmt.Printf("raw str:%s\n", sr)

	z := "中国人"

	for _, v := range z {
		fmt.Printf("unicode:%x\n", v)
	}
	d := '中'
	fmt.Printf("中unicode:%x\n", d)
	fmt.Println("--------------------------")
	for i := 0; i < len(z); i++ {
		fmt.Printf("i:0x%x\n", z[i])
	}
	/**
	utf-8编码原理
	#1-byte characters have the following format: 0xxxxxxx  			: U+0000 -> U+007F
	#2-byte characters have the following format: 110xxxxx 10xxxxxx		        : U+0080 -> U+07FF
	#3-byte characters have the following format: 1110xxxx 10xxxxxx 10xxxxxx	  : U+0800 -> U+FFFF
	#4-byte characters have the following format: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx : U+10000 -> U+10FFFF
	*/
}
