package main

import "fmt"

func main() {
	fmt.Print("Введите число: ")
	var F float64
	fmt.Scanf("%f", &F)

	M := F * 0.3048

	fmt.Println(M)
}
