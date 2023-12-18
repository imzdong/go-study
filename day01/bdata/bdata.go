package bdata

import (
	"fmt"
	"unicode/utf8"
)

// 常量
const (
	c1 = 6
	c2 = 7
)

type Bdata struct {
	Name string
	Age int
}

func (bd *Bdata) Day01DataVar() {
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

func Day01DataNum() {
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

	zh := "中国人"
	//rune切片
	zhr := []rune(zh)
	fmt.Printf("str->rune:%x\n", zhr)
	//字节切片
	zhb := []byte(zh)
	fmt.Printf("str->byte:%b\n", zhb)

	zr := string(zhr)
	fmt.Printf("rune->str:%s\n", zr)

	zb := string(zhb)
	fmt.Printf("byte->str:%s\n", zb)
}

func foo(arr [5]int) {

}
func day01Array() {
	var arr1 [3]int
	//foo(arr1) error
	fmt.Println(arr1)
	var arr2 [5]int
	foo(arr2)
	var arr3 [5]string
	fmt.Println(arr3)
	//foo(arr3) error
	//var arr [N]T

	arr4 := [4]string{"z", "3"}
	fmt.Println(arr4)

	arr5 := [...]string{}
	fmt.Println(len(arr5))

	arr6 := [...]string{
		6: "dong",
		}
		fmt.Println(arr6)

	a0 := [3]string{"1", "2", "dong"}
	fmt.Printf("a0 t:%T\n", a0)

	a := []string{"1", "2", "dong"}
	fmt.Printf("a t:%T\n", a)
	for _, v := range a {
		fmt.Println(v)
	}
	//fmt.Println(a[2])

	//make 构建切片 10底层数组长度
	mq := make([]string, 5, 10)
	fmt.Println(len(mq))

	//采用array[low : high : max]语法基于一个已存在的数组创建切片。这种方式被称为数组的切片化
	arr7 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr7)

	arr8 := arr7[1:3:6]
	//len=3-1 cap=6-1
	fmt.Println(arr8)

}

func day02Map() {
	//函数类型、map类型自身，以及切片类型是不能作为map的key类型的。
	//map[keyType]valueType
	m := map[int]string{1: "dong"}
	fmt.Println(m)
	fmt.Println(m[1])
	v, ok := m[0]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	m1 := make(map[int]string, 6)
	fmt.Println(m1)
	fmt.Println(len(m1))

	s := make([]int, 3)
	fmt.Println(s)
	fmt.Printf("%T", s)

	//make方法可以创建 切片，map，chan，主要是引用类型
	//struct，基本类型，string都是用字面值声明
}

func (bd *Bdata) Day02Struct() {
	type person struct {
		age  int
		name string
	}

	p := person{
		age:  18,
		name: "dong",
		}
		fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println(p.name)
}