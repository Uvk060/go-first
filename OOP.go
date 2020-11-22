package main

import "fmt"

type MyObject struct {
	MyString string
	MyInt    int
}

func main() {
	thatObject := new(MyObject)          //link to the object
	thatObject.MyString = "GO language." //эта строка равна тому-то и в принте это отражается + ссылка на объект
	//thatObject := MyObject{
	//MyString: "Hello"
	//}      //так мы получаем сразу значение объекта, а через new сначала ссылку //this way is BETTER
	fmt.Println(thatObject) //link to the object
}

func (m *MyObject) ChangeString(str string) { //метод к данному типу. но он не меняет исходный тип! только его копию
	fmt.Println(m.MyString, str)
}
