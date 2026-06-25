package main

import (
	"fmt"
)

func main() {
	var year = 2017
	if 1946 <= year && year <= 1964 {
		fmt.Println("Hello old man")
	} else if 1965 <= year && 1980 >= year {
		fmt.Println("Hello X")
	} else if 1981 <= year && year <= 1996 {
		fmt.Println("Hello millenial")
	} else if 1997 <= year && year <= 2012 {
		fmt.Println("Hello Z")
	} else if 2013 <= year {
		fmt.Println("Hello alpha")
	}
}
