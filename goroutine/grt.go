package goroutine

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Concurrently ninja with delay")
	}()
	go fmt.Println("goroutine 1")
	go fmt.Println("goroutine 2")
	go fmt.Println("goroutine 3")

	time.Sleep(2 * time.Second)
	fmt.Println("goroutine 4")
}
