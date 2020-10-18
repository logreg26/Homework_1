package main

import (
	"fmt"
	"math"
)

func main() {

	catA := 10
	catB := 20

	gyp := math.Sqrt(math.Pow(float64(catA), 2) + math.Pow(float64(catB), 2))
	perim := float64(catA) + float64(catB) + float64(gyp)
	square := math.Sqrt((float64(perim) / 2) * (((float64(perim) - float64(catA)) / 2) * ((float64(perim) - float64(catB)) / 2) * ((float64(perim) - float64(gyp)) / 2)))

	fmt.Println("Гипотенуза треугольника = ", gyp)
	fmt.Println("Периметр треугольника = ", perim)
	fmt.Println("Площадь треугольника = ", square)
}
