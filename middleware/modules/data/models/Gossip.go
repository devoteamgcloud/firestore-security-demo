package models

type GossipModel struct {
	Title   string `firestore:"title"`
	Content string `firestore:"content"`
	User    string `firestore:"user"`
	Policy  string `firestore:"policy"`
}
