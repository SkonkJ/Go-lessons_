package main

import "fmt"

func main() {
	var number uint16 = 260
	number = number + 19
	fmt.Println("hello World", number)
	var i int = 15
	var s string = "Canada"
	var a float32 = 1000
	var hello string
	hello = "17, 32, Hello world"
	name1 := "Almaz"

	var (
		g = 1000 - 7
		h = 10
		j = 33
	)
	var (
		age  int    = 15
		name string = "Almaz"
	)

	fmt.Println(hello)
	fmt.Println(g, j, h)
	fmt.Println("I am old", i)
	fmt.Println("I live in", s)
	fmt.Println("I have $", a)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name1)

}
