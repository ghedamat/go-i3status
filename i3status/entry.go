package i3status

import (
	"encoding/json"
)

type Entry struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Button   int    `json:"button"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

func NewEntry(str string) *Entry {
	var e Entry
	json.Unmarshal([]byte(str), &e)
	return &e
}
