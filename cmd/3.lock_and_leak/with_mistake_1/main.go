package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	produceHandler := func(in <-chan int) <-chan int {
		res := make(chan int)

		go func() {
			defer close(res)

			for num := range in {
				fmt.Printf("-> Producer принял число %d\n", num)
				n := num * num
				res <- n
				fmt.Printf("<- Producer отправил число %d\n", n)
			}
		}()

		return res
	}

	consumeHandler := func(produceCh <-chan int) {
		for num := range produceCh {
			fmt.Printf("* Consumer получил число %d\n", num)
		}

		fmt.Println("Я всё получил!")
	}

	in := make(chan int)
	defer close(in)

	produceCh := produceHandler(in)
	consumeHandler(produceCh)

	n := 5
	for i := 0; i <= n; i++ {
		in <- gofakeit.IntRange(1, 10)
	}
}
