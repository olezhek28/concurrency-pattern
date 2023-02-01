package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
)

const (
	Avito    = "avito"
	Cian     = "cian"
	Yandex   = "yandex"
	Domofond = "domofond"
)

type Info struct {
	ServiceName     string
	FlatDescription string
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	services := []string{Avito, Cian, Yandex, Domofond}

	channels := make([]<-chan interface{}, 0, len(services))
	for _, service := range services {
		channels = append(channels, handler(ctx, service))
	}

	resCh := fanIn(ctx, channels...)
	for v := range resCh {
		msg, ok := v.(*Info)
		if !ok {
			continue
		}
		fmt.Println(color.GreenString("Кайфовое предложение от %s: %s", color.RedString(msg.ServiceName), msg.FlatDescription))
	}
}

func handler(ctx context.Context, serviceName string) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		defer close(res)

		for {
			select {
			case <-ctx.Done():
				fmt.Println(color.BlueString("Кто-то выключил %s", serviceName))
				return
			default:
				time.Sleep(time.Second)
				res <- &Info{ServiceName: serviceName, FlatDescription: gofakeit.Address().Address}
			}
		}
	}()

	return res
}
