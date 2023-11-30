package main

func printArrayS(myArray []int) {
	//显然这个是引用传递
	for _, value := range myArray {
		println(value)
	}
	myArray[0] = 100
}

//func main() {
//
//	myArray := []int{1, 2, 3, 4}
//
//	fmt.Printf("myArrat type is %T\n", myArray)
//
//	printArrayS(myArray)
//
//	fmt.Println("================")
//
//	for _, value := range myArray {
//		println(value)
//	}
//}
