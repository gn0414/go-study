package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}

func (this *Human) Walk() {
	fmt.Println("Human walk...")
}

//func main() {
//	h := Human{"zhang3", "female"}
//	h.Eat()
//	h.Walk()
//
//	s := SuperMan{Human{"li4", "female"}, 88}
//	s.Walk()
//	s.Eat()
//	s.fly()
//
//	var s1 SuperMan
//
//	s1.name = "zhao"
//	s1.sex = "male"
//
//	s1.fly()
//}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法 ??我c那不是随便继承
	level int
}

// 重写
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat...")
}

// 子类新方法
func (this *SuperMan) fly() {
	fmt.Println("SuperMan Fly...")
	println(this.name)
	println(this.sex)

}
