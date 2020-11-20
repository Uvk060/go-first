package main

import "fmt"

type OwnInt uint8

func main() {
	//a := OwnInt(15) //указатель
	a := new(OwnInt)
	//указатель на Ownint, но для присвоения значения нужно разъименовать
	*a = 15
	fmt.Println(a)

}
