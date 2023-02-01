package main

import (
	"log"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan struct{}, urls ...string) <-chan *http.Response {
		resCh := make(chan *http.Response)
		go func() {
			defer close(resCh)

			for _, url := range urls {
				res, err := http.Get(url)
				if err != nil {
					log.Printf("Что-то барахлит этот сайт %s: %s", url, err.Error())
					continue
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
	for response := range checkStatus(done, urls...) {
		log.Printf("А вот это сайт %s вернул мне такой код: %v\n", response.Request.Host, response.Status)

		err := response.Body.Close()
		if err != nil {
			log.Printf("У меня не вышло закрыть тело ответа: %s", err.Error())
		}
	}
}
