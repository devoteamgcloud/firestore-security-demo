package entity

import "cloud.google.com/go/firestore"

type User interface {
	Id() string
	Database() *firestore.Client // this isn't good practice, decouple this db
	Role() RoleCode
	Department() DepartmentCode
}
