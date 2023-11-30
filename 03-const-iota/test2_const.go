package main

//定义枚举

const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

const (
	SICHUAN = iota
	GUANGZHOU
	FUJIAN = iota * 2
	JIANGXI
	GUANGYUAN = iota * 3
)

func main() {

	//定义常量

	const a = 100

	const b int = 100

	println(BEIJING)
	println(SHANGHAI)
	println(SHENZHEN)
	println(SICHUAN)
	println(GUANGZHOU)
	println(FUJIAN)
	println(JIANGXI)
	println(GUANGYUAN)
}
