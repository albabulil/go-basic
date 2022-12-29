package main

import "fmt"

func main() {

	//operator matematika
	var value = 10 + 20
	fmt.Println(value)

	var i = 10
	var j = 10
	var k = i * j
	fmt.Println(k)

	//augmented assigments
	var a = 10
	var b = 10
	var c = 20
	a += a + b + c // a = a +10 + 10 + 20
	fmt.Println(a)

	//unary operator
	i++ // i = i + 1
	j-- // i = i - 1
	fmt.Println(i)
	fmt.Println(j)

}
