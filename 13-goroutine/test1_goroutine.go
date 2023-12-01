package main

import (
	"fmt"
	"time"
)

// 子 goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine:i = %d\n", i)
		time.Sleep(1 * time.Second)
	}

}

// 主 goroutine
func main() {
	//创建一个go程去执行newTask() 流程
	//go newTask()
	//
	//i := 0
	//
	//for {
	//	i++
	//	fmt.Printf("main goroutine:i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

	//匿名方法启动携程
	//用go承载一个形参为空返回值为空的一个函数
	//go func() {
	//	defer fmt.Println("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		//退出goroutine
	//		runtime.Goexit()
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()

	//有参数的go程
	//我希望我的主go程能得到我的子go程的一些信息
	//我们只能来通过channel了
	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
