package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	produceHandler := func(done <-chan struct{}, in <-chan int) <-chan int {
		res := make(chan int)

		go func() {
			defer close(res)

			for {
				select {
				case num, ok := <-in:
					if ok {
						fmt.Printf("-> Producer принял число %d\n", num)
						n := num * num
						res <- n
						fmt.Printf("<- Producer отправил число %d\n", n)
					}

				case <-done:
					fmt.Println("Producer: кто-то отменил меня")
					return
				}
			}
		}()

		return res
	}

	consumeHandler := func(done <-chan struct{}, produceCh <-chan int) {
		count := 0
		//defer fmt.Printf("Consumer: я получил %d чисел!\n", count)
		//defer func() {
		//	fmt.Printf("Consumer: я получил %d чисел!\n", count)
		//}()

		for {
			select {
			case num, ok := <-produceCh:
				if ok {
					fmt.Printf("* Consumer получил число %d\n", num)
					count++
				}

			case <-done:
				fmt.Println("Consumer: кто-то отменил меня")
				return // Может break?
			}
		}

		// fmt.Printf("Consumer: я получил %d чисел!\n", count)
	}

	in := make(chan int)
	done := make(chan struct{})

	produceCh := produceHandler(done, in)

	go func() {
		//defer close(in)

		n := 5
		for i := 0; i <= n; i++ {
			in <- gofakeit.IntRange(1, 10)

			//if i == n/2 {
			//	close(done)
			//}
		}
	}()

	consumeHandler(done, produceCh)
}
