package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost   = 12
	minFirstname = 3
	minLastName  = 3
	minPassword  = 7
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	Lastname          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func (params CreateUserParams) Validate() error {
	if len(params.FirstName) < minFirstname {
		return fmt.Errorf("first name should be %d chars long", minFirstname)
	}
	if len(params.Lastname) < minLastName {
		return fmt.Errorf("last name should be %d chars long", minLastName)
	}
	if len(params.Password) < minPassword {
		return fmt.Errorf("passwqrd should be %d chars long", minPassword)
	}
	return nil

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
