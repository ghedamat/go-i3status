package main

import (
	"fmt"
	"github.com/ghedamat/go-i3status/i3status"
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
	b := i3status.NewBar()

	b.Add(i3status.NewTimerWidget())
	b.Add(i3status.NewDateWidget())
	//b.Add(i3status.NewOnOffWidget())
	//w4 := i3status.NewI3statusWidget()
	b.Add(i3status.NewEchoWidget())

	for {
		m := <-b.Output
		fmt.Println(m + ",")
	}

}
