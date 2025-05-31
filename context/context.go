package context

import (
	"context"
	"fmt"
	"time"
)

func contextplain() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
	ctx = context.WithValue(ctx, "id", 1)

	//go func() {
	//	time.Sleep(time.Millisecond * 100)
	//	cancel()
	//}()

	parse(ctx)
}

func parse(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Println(id)
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing complete")
			return
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}
}
