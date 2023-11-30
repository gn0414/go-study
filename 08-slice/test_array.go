package main

// 那就是只能传入10个长度的数组的类型
func printArray(myArray [10]int) {
	//并且这个是值拷贝
	for index, value := range myArray {
		println(index, "-", value)
	}
	//外面的根本不会变
	myArray[0] = 111
}

//func main() {
//	//固定长度的一个数组
//	var myArray1 [10]int
//	for i := 0; i < len(myArray1); i++ {
//		println(myArray1[i])
//	}
//
//	myArray2 := [10]int{1, 2, 3, 4}
//
//	for index, value := range myArray2 {
//		println(index, "-", value)
//	}
//
//	printArray(myArray1)
//
//	fmt.Printf("myArray1 types = %T\n", myArray2)
//}
