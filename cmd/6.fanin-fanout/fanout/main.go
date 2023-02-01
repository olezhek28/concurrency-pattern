package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Decision struct {
	ServiceName string
	IsAgree     bool
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mainChan := make(chan interface{})
	defer close(mainChan)

	channels := fanOut(ctx, mainChan)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, c := range channels {
		c := c
		go func() {
			defer wg.Done()
			for msg := range c {
				decision, ok := msg.(*Decision)
				if !ok {
					continue
				}

				fmt.Println(color.GreenString("Ответ сервису %s: %v", color.RedString(decision.ServiceName), decision.IsAgree))
			}
		}()
	}

	for i := 0; i < 10; i++ {
		mainChan <- &Decision{ServiceName: "avito", IsAgree: true}
	}

	wg.Wait()
}
