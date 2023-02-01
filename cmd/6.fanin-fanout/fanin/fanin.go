package main

import (
	"context"
	"sync"
)

func fanIn(ctx context.Context, channels ...<-chan interface{}) <-chan interface{} {
	wg := sync.WaitGroup{}

	res := make(chan interface{})
	multiplexer := func(channel <-chan interface{}) {
		defer wg.Done()
		for v := range channel {
			select {
			case <-ctx.Done():
				return
			default:
				res <- v
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplexer(c)
	}

	go func() {
		defer close(res)
		wg.Wait()
	}()

	return res
}
