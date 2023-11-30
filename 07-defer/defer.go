package main

import "fmt"

func func1() {
	println("a")
}
func func2() {
	println("b")
}

func func3() {
	println("c")
}

func main() {
	//压栈 所以2先执行
	defer fmt.Println("end1")
	defer fmt.Println("end2")

	println("hello go! 1")
	println("hello go! 2")

	//必然是3最先出栈其次是21
	defer func1()
	defer func2()
	defer func3()
}
