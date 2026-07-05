package main

import (
	"fmt"
)

func RemoveDuplicates(input []string) []string {
	solo := make(map[string]int)
	res := make([]string, 0)
	for _, item := range input {
		if v := solo[item]; v == 0 {
			solo[item]++
			res = append(res, item)
		}
	}
	return res
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemoveDuplicates(input))
}
