package examples

import (
	"fmt"
	"time"
)

/*
ChanDialog output:

Hi chan, "how are you? [1]"
Hello other chan "I'm fine [1]"
Hi chan, "how are you? [2]"
Hello other chan "I'm fine [2]"
Hi chan, "how are you? [3]"
Hello other chan "I'm fine [3]"
Hi chan, "how are you? [4]"
Hello other chan "I'm fine [4]"
Hi chan, "how are you? [5]"
Hello other chan "I'm fine [5]"
*/
func ChanDialog() {
	hiChan := getMessageChan("how are you?")
	imFine := getMessageChan("I'm fine")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Hi chan, %q\n", <-hiChan)
		fmt.Printf("Hello other chan %q\n", <-imFine)
	}
}

func getMessageChan(message string) <-chan string {
	c := make(chan string)
	go func() { // Go routine inside a channel
		for i := 1; ; i++ {
			c <- fmt.Sprintf("%s [%d]", message, i)
			time.Sleep(time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}
