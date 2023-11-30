package main

import (
	"go-study/05-init/lib1"
)

func main() {
	//我们导包就会先去执行包的init方法,如果包里面仍然有包那么依然进去执行
	println("初始化完成")
	lib1.TestLib1()
}
