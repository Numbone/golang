package goroutine

import "fmt"

export func GoRoutine() {
	go fmt.Println("goroutine 1")
	go fmt.Println("goroutine 2")
	go fmt.Println("goroutine 3")
	fmt.Println("goroutine 4")
}
