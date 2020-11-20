package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
	fmt.Println("Hello, Vadim.")
	fmt.Printf("my number is %d", 22) // %d выыодим цифры
	MyFunc()
	MyNumb(52)
}

func MyFunc() {
	fmt.Println("\n", "Func")
}

func MyNumb(abc int) {
	fmt.Println(abc)
}
