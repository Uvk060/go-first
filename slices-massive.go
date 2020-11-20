package main

import "fmt"

func main() {
	//TwiceNum()
	//twice := [] string("Hello", "Name") ????
	//numbers := []float64{126.03, 77.19, 44.53, 45.55}        ??
	// все массивы передаются по ссылке и меняются из любого места changeArray(numbers)
	arr := make([]uint, 5)          // 2-d method массив из 5 элементов
	arr2 := append(arr, 12, 13, 14) //добавила 3 значения в массив
	//fmt.Println(len(arr2))          //вывела длину массива = 8
	fmt.Println(arr2[len(arr2)-1]) //print the last element of the array
	//for item := 0; item<len(arr2); item++{
	//	fmt.Println(arr2[item]) //прошлись по всем элмементам массива

	//}
	//for key, value := range arr2 {
	//	fmt.Println("key is", key)
	//	fmt.Println("Value is", value) //2 method to run to the array
	//}

	//for _, value := range arr2 {
	//	fmt.Println("Value is", value) //3 method to run to the array, if not important key
	//}
}

//func TwiceNum(string){                      ????
//twice := [] string("Hello", "Name")
//fmt.Println(twice)
//}
