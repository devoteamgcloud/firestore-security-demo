package models

type UserModel struct {
	Role string `firestore:"role"`
}
