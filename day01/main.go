package main

import (
	"fmt"
	"github.com/imzdong/day01/bdata"
)

func main() {
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
