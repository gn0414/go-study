package main

func main() {

}

// 函数正常定义
func test(a string) int {
	return 1
}

// 多返回值(无名返回值)
func test1(b string) (int, int) {
	return 2, 3
}

// 有名返回值
func test2(c string) (r1 int, r2 int) {
	r1 = 1
	r2 = 1
	return
}

// 返回值类型相同
func test3(d string) (r1, r2 int) {
	return 3, 4
}
