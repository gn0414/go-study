package main

import "fmt"

// 我们的多态就要通过接口来实现
// 本质是一个指针
type AnimaIf interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	Color string
	Type  string
}

//实现这些接口方法其实就是实现了这个接口
//必须要全部实现！

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) GetType() string {
	return this.Type
}

type Dog struct {
	Color string
	Type  string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string {
	return this.Color
}
func (this *Dog) GetType() string {
	return this.Type
}
func main() {
	var animal AnimaIf
	//所以说为什么说是指针这下懂了吧牢底
	animal = &Cat{"Green", "coke"}

	animal.Sleep()

	var animal2 AnimaIf
	animal2 = &Dog{"jack", "jinmao"}

	animal2.Sleep()
	//Sleep多态了牢底汗流浃背了吧

	book := Book{auth: "simon"}
	myFunc(book)

	myFunc(100)

	myFunc("abc")

	myFunc(3.14)
}

//interface{}是万能数据类型

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	//interface{}如何区分引用的底层数据类型到底是什么？
	//断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type", value)
	}
}

type Book struct {
	auth string
}
