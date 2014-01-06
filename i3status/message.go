package i3status

import (
	"encoding/json"
	"log"
)

type Message struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Color               string `json:"color"`
	MinWidth            int    `json:"min_width"`
	Align               string `json:"align"`
	Name                string `json:"name"`
	Instance            string `json:"instance"`
	Urgent              bool   `json:"urgent"`
	Separator           bool   `json:"separator"`
	SeparatorBlockWidth int    `json:"separator_block_width"`
}

func (m *Message) ToJson() string {
	s, err := json.Marshal(m)
	if err != nil {
		log.Fatal("failed to encode message")
	}

	return string(s)
}

func NewMessage() *Message {
	return &Message{
		Separator:           true,
		Align:               "left",
		SeparatorBlockWidth: 10,
	}
}
