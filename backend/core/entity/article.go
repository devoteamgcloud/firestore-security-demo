package entity

type Article interface {
	Title() string
	Content() string
	User() string
	//Policy() PolicyCode
}
