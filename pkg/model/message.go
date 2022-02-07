package model

type Message struct {
	Client string  `json:"client"`
	Files  []*File `json:"files"`
}
