package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string, value:"aceld">
	a = "aceld"

	//pair<type:string, value:"aceld">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

	//我们的变量由type和value的一个二元组组成
	// b: pair<type:Book,value:book{}的地址>
	b := &Book1{}
	//:pair<type:,value>
	var r Reader
	// r: pair<type:Book,value:book{}的地址>
	//指针pair就指向了我们的变量的type 和 value
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	//为什么断言能成功呢？ 因为确实是w和r的指针指向的type都是一致的
	w.WriteBook()

}

type Book1 struct {
}
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

func (this *Book1) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book1) WriteBook() {
	fmt.Println("write a book")
}
