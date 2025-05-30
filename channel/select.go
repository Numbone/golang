package channel

import (
	"fmt"
	"time"
)

func selectel() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "message 1 spent 200 ms"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for {
			message2 <- "message 2 spent 1 s"
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select { // not blocked
		case msg1 := <-message1:
			fmt.Println(msg1)
		case msg2 := <-message2:
			fmt.Println(msg2)
		}
		//fmt.Println(<-message1)
		//fmt.Println(<-message2)
	}
}
