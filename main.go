package main

import "fmt"

func main() {
	fmt.Println(len("Я люблю аниме и не думаю о том, что другие люди считаю по этому поводу"))
	fmt.Println("Sалава плюс поплава"[0])
	fmt.Println("ТвояМать" + "Шлюха")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(false && false)
	var number int64
	fmt.Scanf("%d\n", &number)
	var number2 int64
	var number3 int64
	var number4 int64
	fmt.Scanf("%d\n", &number2)
	fmt.Scanf("%d\n", &number3)
	fmt.Scanf("%d\n", &number4)
	fmt.Println(number + number2)
	fmt.Println(number3 * number4)
	fmt.Println(number / number2)
}
