package main

import "fmt"

func main() {

	//第一种声明
	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d,slice = %v\n", len(slice1), slice1)

	//第二种声明
	//位slice开辟三个空间,默认值是0

	var slice2 = make([]int, 3)
	slice2[0] = 100
	fmt.Printf("len = %d,slice = %v\n", len(slice2), slice2)

	//第三种
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d,slice = %v\n", len(slice3), slice3)

	//第四种
	slice4 := make([]int, 3)
	fmt.Printf("len = %d,slice = %v\n", len(slice4), slice4)

	//判断一个slice是否位空
	var slice5 []int
	if slice5 == nil {
		println("slice5 is nil")
	}

}
