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

func fanOut(ctx context.Context, mainChan <-chan interface{}) []chan interface{} {
	services := []string{Avito, Cian, Yandex, Domofond}

	channelMap := make(map[string]chan interface{}, len(services))
	for _, service := range services {
		channelMap[service] = handler(ctx, service)
	}

	go func() {
		for v := range mainChan {
			msg, ok := v.(Decision)
			if !ok {
				continue
			}

			switch msg.ServiceName {
			case Avito:
				channelMap[Avito] <- msg
			case Cian:
				channelMap[Cian] <- msg
			case Yandex:
				channelMap[Yandex] <- msg
			case Domofond:
				channelMap[Domofond] <- msg
			}
		}
	}()

	channels := make([]chan interface{}, 0, len(services))
	for _, ch := range channelMap {
		channels = append(channels, ch)
	}

	return channels
}

func handler(ctx context.Context, serviceName string) chan interface{} {
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
				res <- &Decision{ServiceName: serviceName, IsAgree: gofakeit.Bool()}
			}
		}
	}()

	return res
}
