package users_manager

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserCreateInput struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `json:"name" validate:"required,min=3,max=50"`
	Email    string        `json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=8"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateNameInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	NewName  string `json:"new_name" validate:"required,min=3,max=50"`
}

type UserUpdateEmailInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	NewEmail string `json:"new_email" validate:"required,email"`
}

func Validate(u interface{}) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
