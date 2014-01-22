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

	b.Add(i3status.NewTimerWidget())
	//b.Add(i3status.NewDateWidget())
	//b.Add(i3status.NewOnOffWidget())
	//w4 := i3status.NewI3statusWidget()
	b.Add(i3status.NewEchoWidget())

	for {
		fmt.Println(b.Message() + ",")
		time.Sleep(1 * 1e9)
	}

}
