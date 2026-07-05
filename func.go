package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(fig figures) (func(float64) float64, bool) {
	switch fig {
	case square:
		return func(p float64) float64 { return p * p }, true
	case circle:
		return func(f float64) float64 { return math.Pi * f * f }, true
	case triangle:
		return func(f float64) float64 { return math.Sqrt(3 * f * f * f * f / 16.0) }, true
	default:
		return nil, false
	}
}

func main() {
	myFigure := square
	var x float64 = 2.0
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(x)
	fmt.Println(myArea)
}
