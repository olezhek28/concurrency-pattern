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
	eat := func(people []*Properties) []*Properties {
		res := make([]*Properties, 0, len(people))
		for _, v := range people {
			v.IsEat = gofakeit.Bool()
			res = append(res, v)
		}

		return res
	}

	putOnHat := func(people []*Properties) []*Properties {
		res := make([]*Properties, 0, len(people))
		for _, v := range people {
			v.IsPut = gofakeit.Bool()
			res = append(res, v)
		}

		return res
	}

	hearAboutMothersFriendSon := func(people []*Properties) []*Properties {
		res := make([]*Properties, 0, len(people))
		for _, v := range people {
			v.IsHear = gofakeit.Bool()
			res = append(res, v)
		}

		return res
	}

	n := 5
	people := make([]*Properties, 0, n)
	for i := 0; i < n; i++ {
		people = append(people, &Properties{Name: gofakeit.Name()})
	}

	start := time.Now()
	people = hearAboutMothersFriendSon(putOnHat(eat(people)))
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
