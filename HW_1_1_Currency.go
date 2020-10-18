package main

import "fmt"

func main() {
	fmt.Print("Введите количество рублей, которые Вы хотите конвертировать: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input / 80

	fmt.Println(output)
}
