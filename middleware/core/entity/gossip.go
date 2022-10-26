package entity

type Gossip interface {
	Title() string
	Content() string
	User() string
	//Policy() PolicyCode
}
