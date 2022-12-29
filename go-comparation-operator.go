package main

import "fmt"

func main() {

	//it's same
	var name1 = "ulil"
	var name2 = "ulil"
	var result bool = name1 == name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 200
	fmt.Println(val1 < val2)
	fmt.Println(val1 > val2)
	fmt.Println(val1 >= val2)
	fmt.Println(val1 <= val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)

}
