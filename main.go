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
	fmt.Println(`{"version":1,"click_events": true}`)
	fmt.Println("[")
	c := make(chan i3status.Message)
	b := i3status.NewBar(c)

	w2 := i3status.NewDateWidget()
	w3 := i3status.NewOnOffWidget()
	w4 := i3status.NewI3statusWidget()
	b.Add(w2)
	b.Add(w3)
	b.Add(w4)

	for {
		fmt.Println(b.Message() + ",")
		time.Sleep(1 * 1e9)
	}

}
