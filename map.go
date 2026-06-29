package main

import "fmt"

func main() {
	prices := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	for k, v := range prices {
		if v > 500 {
			fmt.Print(k, " ")
		}
	}
	fmt.Println()
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	sum := 0
	for _, elem := range order {
		sum += prices[elem]
	}
	fmt.Println(sum)
}
