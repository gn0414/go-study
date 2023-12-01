package main

import "fmt"

func changeValue(slice []int) {
	slice[3] = 10
}

func main() {

	//我们的切片make时第二个参数表示长度,第三个参数表示容量
	slice1 := make([]int, 3, 5)
	slice1[0] = 100
	fmt.Printf("slice =  %v len = %v cap= %v", slice1, len(slice1), cap(slice1))

	//如果我们要追加slice的内容可以使用append
	slice1 = append(slice1, 2)
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 3)
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 4)
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice1, len(slice1), cap(slice1))

	//我们运行时可以观察到,当我们的slice的len超过了我们的cap,我们的cap就会进行两倍的扩容（若是这个slice足够大,我们扩容机制会变为1/4）
	//我们slice的截取一般为
	slice2 := slice1[1:2]
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice2, len(slice2), cap(slice2))
	//可以看到我们的下标从1开始，不包括2所以只有一个,且cap(因为我们的首指针为1 所以是10-1)
	slice3 := slice1[2:]
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice3, len(slice3), cap(slice3))
	//这个就是从2到最后一个了

	slice4 := slice1[:5]
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice4, len(slice4), cap(slice4))
	//这个就是从0到4 0包含了

	//可以看到我们的slice作为参数传递是指针（即引用传递）
	changeValue(slice1)
	println()
	fmt.Printf("slice =  %v len = %v cap= %v", slice1, len(slice1), cap(slice1))

	//copy可以将底层数组的slice一起拷贝

	s1 := []int{1, 2, 3}

	s2 := make([]int, 3) //s2 = [0 0 0]

	copy(s2, s1)

	//将s1中的值依次拷贝到s2
	//值肯定是一样的一看这个就是值拷贝了
	println()
	println(s1)
	println(s2)

	//截取就不是值拷贝了
	//举个例子

	s3 := s1[:]
	println(s3)

	//可以看到头指针都是一样的

}
