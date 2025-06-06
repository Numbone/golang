package channel

import "fmt"

func main() {
	msg := make(chan string, 3) // buffer -->deadlock

	msg <- "channel"
	msg <- "channel 2"
	msg <- "channel 3"

	//go func() {
	//	time.Sleep(time.Second * 1)
	//	msg <- "channel"
	//	msg <- "channel 2"
	//	msg <- "channel 3"
	//}()

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	close(msg)

	//not work if assign goroutine
	for m := range msg {
		fmt.Println(m)
	}
}
