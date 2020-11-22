package main

import "fmt"

type MyObject struct {
	MyString string
	MyInt    int
}

func main() {
	thatObject := new(MyObject)          //link to the object
	thatObject.MyString = "GO language." //эта строка равна тому то и в принте это отражается + ссылка на объект
	fmt.Println(thatObject)              //link to the object
}
