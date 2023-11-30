package main

func main() {
	returnAndDefer()
}

func deferFunc() int {
	println("defer func called...")
	return 0
}

func returnFunc() int {
	println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()

	//return是会比defer先返回 我们就清楚我们的defer处于函数的生命周期完结之后
}
