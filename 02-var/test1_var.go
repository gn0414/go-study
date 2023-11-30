package main

import "fmt"

// 全局变量只能用方式一二三来写第四不行
var aa int
var bb int = 100
var cc = 200

func main() {
	//方法一: 声明一个变量 默认值为0
	var a int
	fmt.Println(a)
	//方式二
	var b int = 100
	println(b)
	//方式三
	var c = 200
	println(c)
	//方式四
	d := 300
	println(d)

	println(aa)
	println(bb)
	println(cc)
}
