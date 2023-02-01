package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	n := 10

	produceHandler := func() <-chan int {
		res := make(chan int, n)

		go func() {
			defer close(res)

			for i := 0; i <= n; i++ {
				res <- gofakeit.IntRange(1, 100)
			}
		}()

		return res
	}

	consumeHandler := func(produceCh <-chan int) {
		for num := range produceCh {
			fmt.Println(num)
		}

		fmt.Println("Я всё получил!")
	}

	produceCh := produceHandler()
	//close(produceCh) // Мы не можем закрыть read-only канал и это нас оберегает от ошибок
	consumeHandler(produceCh)
}
