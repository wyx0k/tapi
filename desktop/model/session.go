package model

import "encoding/json"

type Session struct {
	TabSession TabSession `json:"tab_session"`
}

func (s *Session) Marshal() []byte {
	bytes, _ := json.Marshal(s)
	return bytes
}

func (s *Session) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}

type TabSession struct {
	Tabs          []Tab  `json:"tabs"`
	TabScrollLeft string `json:"tab_scroll_left"`
	CurrentTab    Tab    `json:"current_tab"`
}
type TabType string

const (
	TabRequest    TabType = "request"
	TabCollection TabType = "collection"
)

type Tab struct {
	Type TabType `json:"type"`
	Name string  `json:"name"`
	Id   int     `json:"id"`
}
