package main

import (
	"context"
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
	eat := func(ctx context.Context, intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)
			defer fmt.Println(color.BlueString("Столовая закрылась"))

			for {
				select {
				case v, ok := <-intStream:
					if !ok {
						return
					}
					v.IsEat = gofakeit.Bool()
					res <- v
				case <-ctx.Done():
					return
				}
			}
		}()

		return res
	}

	putOnHat := func(ctx context.Context, intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)
			defer fmt.Println(color.BlueString("Шапки закончились"))

			for {
				select {
				case v, ok := <-intStream:
					if !ok {
						return
					}
					v.IsPut = gofakeit.Bool()
					res <- v
				case <-ctx.Done():
					return
				}
			}
		}()

		return res
	}

	hearAboutMothersFriendSon := func(ctx context.Context, intStream <-chan *Properties) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)
			defer fmt.Println(color.BlueString("Я больше не могу говорить о друзьях мамы"))

			for {
				select {
				case v, ok := <-intStream:
					if !ok {
						return
					}
					v.IsHear = gofakeit.Bool()
					res <- v
				case <-ctx.Done():
					return
				}
			}
		}()

		return res
	}

	generator := func(ctx context.Context, n int) <-chan *Properties {
		res := make(chan *Properties)
		go func() {
			defer close(res)
			defer fmt.Println(color.BlueString("Я больше не могу генерировать людей"))

			for i := 0; i < n; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					res <- &Properties{Name: gofakeit.Name()}
				}
			}
		}()

		return res
	}

	n := 10_000
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	start := time.Now()
	resCh := hearAboutMothersFriendSon(ctx, putOnHat(ctx, eat(ctx, generator(ctx, n))))

	people := make([]*Properties, 0, n)
	for v := range resCh {
		people = append(people, v)
	}
	end := time.Now()

	fmt.Println(color.YellowString("Время выполнения для обработки %d человек составило: %v", len(people), end.Sub(start)))
}
