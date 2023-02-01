package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	n := 10

	// Барни готовит орехи
	nuts := make([]int, 0, n)
	for len(nuts) < n {
		nuts = append(nuts, gofakeit.IntRange(1, 100))
	}

	barny := func(portal chan<- int) {
		defer close(portal)

		// Барни отправляет орехи в канал для гуся
		for i := range nuts {
			portal <- nuts[i]
		}
	}

	// Портал между мирами, который охраняет гофер
	portal := make(chan int)
	go barny(portal)

	close(portal) // Потенциально возможная ситуация, когда портал закрывает гусь до того, как Барни успел передать все орехи

	// Гусь собирает орехи из канала и кайфует
	for num := range portal {
		fmt.Printf("Гусь: я съел %d орех(а,ов)\n", num)
	}
}
