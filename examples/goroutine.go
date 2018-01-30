package examples

import (
	"fmt"
)

/*
Goroutine output:

<Nothing>
*/
func Goroutine() {
	go sayHi(5, "Goroutine")
}

func sayHi(times int, message string) {
	for i := 0; i < times; i++ {
		fmt.Sprintf("%s [%d]", message, i)
	}
}
