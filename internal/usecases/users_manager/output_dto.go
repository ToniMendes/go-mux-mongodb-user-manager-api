package users_manager

import "go.mongodb.org/mongo-driver/v2/bson"

type UserCreateResponse struct {
	Name  string `bson:"name" `
	Email string `bson:"email" `
}

type UserGetAllResponse struct {
	Id    bson.ObjectID `bson:"_id,omitempty"`
	Name  string
	Email string
}

type UserGetByEmailResponse struct {
	Id   bson.ObjectID `bson:"_id,omitempty"`
	Name string
}

type UserUpdateNameResponse struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Email    string
	password string
	NewName  string
}
type UserUpdateEmailResponse struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Email    string
	password string
	NewEmail string
}
