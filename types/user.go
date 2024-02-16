package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"_firstName" json:"firstName"`
	Lastname          string             `bson:"_lastName" json:"lastName"`
	Email             string             `bson:"_email" json:"email"`
	EncryptedPassword string             `bson:"_encryptedPassword" json:"-"`
}

func NewUserParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		Lastname:          params.Lastname,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
