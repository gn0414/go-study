package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", this)
}

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

func DoFileAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)

	fmt.Println("input type is:", inputType)

	inputValue := reflect.ValueOf(input)

	fmt.Println("input value is:", inputValue)
	//属性
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}
	//方法
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s:%v", method.Name, method.Type)
	}
}

//func main() {
//	var num float64 = 1.2345
//	reflectNum(num)
//
//	user := User{1, "simon", 18}
//
//	DoFileAndMethod(user)
//
//}
