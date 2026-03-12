package domain

import (
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id           bson.ObjectID `bson:"_id,omitempty"`
	Name         string
	NewName      string
	Email        string
	NewEmail     string
	Password     string
	NewPassword  string
	PasswordHash string
}

func NewUpdateEmail(email, password, newEmail string) *User {
	if len(password) < 8 {
		return nil
	}

	if !strings.Contains(email, "@") {
		return nil
	}

	if email == "" || password == "" || newEmail == "" {
		return nil
	}

	return &User{
		Email:    email,
		Password: password,
		NewEmail: newEmail,
	}
}

func NewUpdateName(email, password, newName string) *User {
	if email == "" || password == "" || newName == "" {
		return nil
	}

	if len(newName) < 3 || len(newName) > 50 {
		return nil
	}

	if len(password) < 8 {
		return nil
	}

	if !strings.Contains(email, "@") {
		return nil
	}

	return &User{
		Email:    email,
		Password: password,
		NewName:  newName,
	}
}

func NewLoginUser(email, password string) *User {
	if (email == "" && !strings.Contains(email, "@")) || password == "" {
		return nil
	}

	if len(password) < 8 {
		return nil
	}

	return &User{
		Email:    email,
		Password: password,
	}
}

func NewUser(name, email, password string) *User {
	if name == "" || email == "" || password == "" {
		return nil
	}

	if len(name) < 3 || len(name) > 50 {
		return nil
	}

	if len(password) < 8 {
		return nil
	}

	if !strings.Contains(email, "@") {
		return nil
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
