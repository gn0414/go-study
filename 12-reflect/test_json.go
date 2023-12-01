package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	//后面tag就是json的标题
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"西游记", 2000, 10, []string{"sunwukong", "zhubajie"}}

	//编码的过程（结构体转化为json）

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码
	//jsonStr = {"title":"西游记","year":2000,"rmb":10,"actors":["sunwukong","zhubajie"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal")
		return
	}

	fmt.Println("movie = ", myMovie)

}
