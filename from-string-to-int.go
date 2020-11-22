package main

import (
	"fmt"
	"strconv"
)

func main() {
	StrItoa()
	StrAtoi()
	StrFloat()
}

func StrItoa() {
	num := strconv.Itoa(23)
	fmt.Println(num)
}

func StrAtoi() {
	num, err := strconv.Atoi("10") //в кач аргумента "bnm" и будет ошибка
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(num)
}
func StrFloat() {
	num, _ := strconv.ParseFloat("32,64", 64)
	fmt.Println(num)

}
