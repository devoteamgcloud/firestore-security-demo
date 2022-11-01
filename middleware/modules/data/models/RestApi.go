package models

type RestValue struct {
	Value string `json:"stringValue"` // https://cloud.google.com/firestore/docs/reference/rest/Shared.Types/ArrayValue#Value
}

type RestArticleMap struct {
	Content RestValue `json:"content"`
	Title   RestValue `json:"title"`
}

type RestDoc struct {
	Fields RestArticleMap `json:"fields"`
}

type RestList struct {
	Docs []RestDoc `json:"documents"`
}
