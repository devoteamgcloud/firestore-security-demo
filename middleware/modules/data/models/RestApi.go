package models

type RestValue struct {
	Value string `json:"stringValue"`
}

type Rest struct {
	Gossip RestValue `json:"content"`
	Title  RestValue `json:"title"`
}

type RestDoc struct {
	Fields Rest `json:"fields"`
}

type RestList struct {
	Docs []RestDoc `json:"documents"`
}
