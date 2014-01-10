package i3status

import (
	"fmt"
	"io"
)

type Subscriber struct {
	subs [](chan Entry)
	In   io.Reader
}

func (s *Subscriber) Subscribe(c chan Entry) {
	s.subs = append(s.subs, c)
}

func (s *Subscriber) Len() int {
	return len(s.subs)
}

func (s *Subscriber) Start() {
	go func() {
		var i string
		for {
			fmt.Fscanf(s.In, "%s", &i)
			for _, c := range s.subs {
				c <- *NewEntry(i)
			}
		}
	}()
}
