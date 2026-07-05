package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	freqs := make(map[int]int)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		freqs[tmp] = i
	}
	fmt.Scan(&s)
	for k, v := range freqs {
		another, ok := freqs[s-k]
		if ok {
			fmt.Println(v, another)
			break
		}
	}
}
