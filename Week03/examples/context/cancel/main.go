// cancel context
package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 该函数返回一个chan
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)

		fmt.Println("gen start")
		wg.Add(1)
		n := 1
		// 该goruntines生成一个一个连续递增的数字
		go func() {
			fmt.Println("gorun ...")
			for {
				select {
				case <-ctx.Done():
					fmt.Println("ctx.Done()")
					wg.Done()
					return
				case dst <- n:
					n++
					fmt.Println("n:", n)
				}
			}
		}()
		fmt.Println("gen end")
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 6 {
			break
		}
	}
	// 因为要打印 ctx.Done() 所以不能放在defer
	// 放在defer wg.Wait 要等待 <-ctx.Done() <-ctx.Done() 又要等待wg.Wait
	cancel()

	wg.Wait()
}
