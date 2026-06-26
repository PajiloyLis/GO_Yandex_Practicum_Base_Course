package main

import "fmt"

func main() {
	a := make([]int, 100, 100)
	for i := 1; i <= 100; i++ {
		a[i-1] = i
	}
	fmt.Println(a)
	a = append(a[:10], a[len(a)-10:]...)
	fmt.Println(a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}
