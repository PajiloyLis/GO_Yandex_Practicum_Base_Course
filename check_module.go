package main

import (
	"fmt"

	myModule "github.com/PajiloyLis/GO_Yandex_Practicum_Base_Course/tree/master/module"
)

func main() {
	if sum := myModule.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
