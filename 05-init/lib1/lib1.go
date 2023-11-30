package lib1

import (
	_ "go-study/05-init/lib2"
	//mylib "go-study/05-init/lib2"
	//. "go-study/05-init/lib2"
)

func TestLib1() {

	//那么如果我们想要导入但不使用包的内容捏？
	//匿名导入
	//也可以别名
	println("lib1 test")
	//mylib.TestLib2()
	//TestLib2() . 就表示不需要包名.api调（最好别这样用,怕有一样的方法冲突）
}

func init() {
	println("lib1 init...")
}
