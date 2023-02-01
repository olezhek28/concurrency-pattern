package main

import "fmt"

func main() {
	in := make(chan int)

	go func() {
		close(in)

		//for i := 0; i <= 5; i++ {
		//	in <- i
		//}
	}()

	for {
		select {
		case num, ok := <-in:
			if ok {
				fmt.Printf("-> Producer принял число %d\n", num)
			} else {
				fmt.Println("Канал закрыт")
				return
			}
		}
	}
}
