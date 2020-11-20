package main

import "fmt"

func main() {
	MyFunc()
	MyNumb(52)
	MyNumbers(17, 226)
	fmt.Println(MyReturn(93, 500))
	fmt.Println(MyBag(23.15, 102.7))
	fmt.Println(MyReturnfloat(333.15, 12.07))
}

func MyFunc() {
	fmt.Println("Func")
}

func MyNumb(abc int) {
	fmt.Println(abc)
}

func MyNumbers(bag, eag int) {
	fmt.Println(bag, eag)
}

func MyReturn(bad, good int) int {
	return bad + good
}

func MyBag(raw, goof float32) (float32, float32) {
	return raw, goof
}

func MyReturnfloat(firstsum, secondsum float64) float64 {
	return firstsum + secondsum
}
