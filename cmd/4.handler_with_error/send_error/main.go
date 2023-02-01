package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Result struct {
	Response *http.Response
	Err      error
}

func main() {
	checkStatus := func(done <-chan struct{}, urls ...string) <-chan *Result {
		resCh := make(chan *Result)
		go func() {
			defer close(resCh)

			for _, url := range urls {
				var res *Result

				urlRes, err := http.Get(url)
				if err != nil {
					res = &Result{Response: nil, Err: errors.Wrap(err, fmt.Sprintf("ошибка при запросе к сайту %s", url))}
				} else {
					res = &Result{Response: urlRes, Err: nil}
				}

				select {
				case <-done:
					return
				case resCh <- res:
				}
			}
		}()

		return resCh
	}

	done := make(chan struct{})
	defer close(done)

	urls := []string{"https://olezhek28", "https://www.ozon.ru/", "https://get-offer"}
	for result := range checkStatus(done, urls...) {
		if result.Err != nil {
			log.Println(result.Err.Error())
			continue
		}

		log.Printf("А вот это сайт %s вернул мне такой код: %v\n", result.Response.Request.Host, result.Response.Status)

		err := result.Response.Body.Close()
		if err != nil {
			log.Printf("У меня не вышло закрыть тело ответа: %s", err.Error())
		}
	}
}
