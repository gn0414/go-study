package main

func swap(pa, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	println(a, b)
}
