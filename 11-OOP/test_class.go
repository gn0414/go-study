package main

import "fmt"

//在go里面是没有class这个东西的

// 我们说过大写的方法可以在其他地方访问
// 那么我们类名首字母大写,表示其他包也能访问
type Hero struct {
	//如果说类的属性首字母大写,表示该属性能对外访问的（对外是指其他的包）
	Name  string
	Ad    int
	Level int
}

//func (this Hero) Show() {
//	fmt.Println("hero = ", this)
//}
//
//func (this Hero) GetName() string {
//	return this.Name
//}
//
//func (this Hero) SetName(newName string) {
//	//this 是调用该方法的对象的一个副本拷贝
//	this.Name = newName
//}

func (this *Hero) Show() {
	fmt.Println("hero = ", this)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//这样才是对当前对象的一个访问
	this.Name = newName
}

//func main() {
//	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
//	hero.Show()
//
//	hero.SetName("lisi")
//
//	hero.Show()
//}
