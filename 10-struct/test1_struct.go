package main

// 声明一种新的数据类型
type myint int

type Book struct {
	title string
	auth  string
}

func printBook(book Book) {
	//值传递
	book.auth = "666"
}

func printBook1(book *Book) {
	book.auth = "777"
}

//
//func main() {
//	var a myint = 10
//	fmt.Println(a)
//	fmt.Printf("%T\n", a)
//
//	var book1 Book
//	book1.title = "golang"
//	book1.auth = "zhang3"
//
//	fmt.Printf("%v \n", book1)
//
//	printBook(book1)
//
//	fmt.Printf("%v \n", book1)
//
//	printBook1(&book1)
//
//	fmt.Printf("%v \n", book1)
//
//}
