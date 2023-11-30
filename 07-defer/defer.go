package main

import "fmt"

func main() {
	//压栈 所以2先执行
	defer fmt.Println("end1")
	defer fmt.Println("end2")

	println("hello go! 1")
	println("hello go! 2")
}
