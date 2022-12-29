package main

import "fmt"

func main() {
	type NoKTP string
	type nikah bool

	var noKTPal NoKTP = "9937487377347438834"
	var nikahAl nikah = true
	fmt.Println(noKTPal)
	fmt.Println(nikahAl)

}
