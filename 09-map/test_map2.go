package main

import "fmt"

func changeValue(myMap map[string]string) {

	//同样传参是引用传递
	myMap["Japan"] = "baga!"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "beijing"

	cityMap["Japan"] = "tokyo"

	cityMap["Usa"] = "new york"

	//遍历

	for k, v := range cityMap {
		fmt.Println(k, "-", v)
	}

	changeValue(cityMap)

	for k, v := range cityMap {
		fmt.Println(k, "-", v)
	}

	//删除
	delete(cityMap, "Japan")

	for k, v := range cityMap {
		fmt.Println(k, "-", v)
	}

	//修改和赋值一样

	cityMap["China"] = "Beijing"
	for k, v := range cityMap {
		fmt.Println(k, "-", v)
	}

}
