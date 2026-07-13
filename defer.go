package main

import "fmt"

var Global = 5

func f() {
	defer func() {
		Global = 5
	}()
	Global = 3
	fmt.Println(Global)

}

func main() {
	f()
	fmt.Println(Global)
}
