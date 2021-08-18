package main

import "fmt"

func main() {
	var x string
	var y string
	y = "солнечный"
	x = "день"
	fmt.Println(x)
	x += " хороший"
	fmt.Println(x)
	x += " " + y
	fmt.Println(y)
	fmt.Println(x)

	const z string = "я люблю гулять"
	fmt.Println(z)

}
