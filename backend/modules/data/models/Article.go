package models

type ArticleModel struct {
	Title   string `firestore:"title"`
	Content string `firestore:"content"`
	User    string `firestore:"user"`
}
