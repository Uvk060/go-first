package main

import "fmt"

const (
	_ = iota      // если хотим пропустить a=0, то нижний_
	b = iota * 10 //если хотим увеличить константы десятками *10, четные *2
	c
	d
	e
	f
)

func main() {
	fmt.Println(c, f, "\n")
	fmt.Println(b, e, d)

}
