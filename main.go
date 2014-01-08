package main

import (
	"fmt"
	"github.com/ghedamat/go-i3status/i3status"
	"time"
)

func main() {

	/*
		go func() {
			for {
				fmt.Println("{\"name\":\"testiii\",\"full_text\":\"ciao\"}")
				time.Sleep(1 * time.Second)
			}
		}()

		go func() {
			var i int
			for {
				fmt.Scanf("%d", &i)
				fmt.Println(i)
			}
		}()
	*/
	fmt.Println(`{"version":1}`)
	fmt.Println("[")
	c := make(chan i3status.Message)
	b := i3status.NewBar(c)

	w1 := i3status.NewBaseWidget(c)
	w2 := i3status.NewDateWidget(c)
	w1.Start()
	w2.Start()
	b.Start()

	for {
		fmt.Println(b.Message() + ",")
		time.Sleep(1 * 1e9)
	}

}
