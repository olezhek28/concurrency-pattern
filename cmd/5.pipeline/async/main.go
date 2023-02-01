package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
)

type Properties struct {
	Name   string
	IsEat  bool
	IsPut  bool
	IsHear bool
}

func main() {
	eat := func(intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)

			for v := range intStream {
				v.IsEat = gofakeit.Bool()
				res <- v
			}
		}()

		return res
	}

	putOnHat := func(intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)

			for v := range intStream {
				v.IsPut = gofakeit.Bool()
				res <- v
			}
		}()

		return res
	}

	hearAboutMothersFriendSon := func(intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)

			for v := range intStream {
				v.IsHear = gofakeit.Bool()
				res <- v
			}
		}()

		return res
	}

	generator := func(n int) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)

			for i := 0; i < n; i++ {
				res <- &Properties{Name: gofakeit.Name()}
			}
		}()

		return res
	}

	n := 5
	start := time.Now()
	resCh := hearAboutMothersFriendSon(putOnHat(eat(generator(n))))

	people := make([]*Properties, 0, n)
	for v := range resCh {
		people = append(people, v)
	}
	end := time.Now()

	printPeople(people)
	fmt.Println(color.YellowString("Время выполнения для обработки %d человек составило: %v", len(people), end.Sub(start)))
}

func printPeople(people []*Properties) {
	for _, v := range people {
		fmt.Printf("%s:\n", v.Name)

		if v.IsEat {
			fmt.Println(color.GreenString("* Покушал(а)"))
		} else {
			fmt.Println(color.RedString("* Ходит голодный(ая)"))
		}
		if v.IsPut {
			fmt.Println(color.GreenString("* Надел(а) шапку"))
		} else {
			fmt.Println(color.RedString("* Ушёл(а) без шапки"))
		}
		if v.IsHear {
			fmt.Println(color.GreenString("* Послушал(а) о сыне маминой подруги"))
		} else {
			fmt.Println(color.RedString("* Кажется у его (её) мамы нет подруг"))
		}
	}
}
