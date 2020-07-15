package domain

import "github.com/getclasslabs/user/internal/pkg"

type User struct {
	//Must Have
	Email    string `validate:"required,required"`
	Name     string `validate:"required,required"`
	PlainPassword string `validate:"required,required"`

	//Generated
	Id int64
	Password string
}

func NewUser(email, name, password string) (*User, error) {
	u := User{
		Email: email,
		Name: name,
		PlainPassword: password,
	}
	err := u.cryptPassword()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) cryptPassword() error {
	password, err := pkg.Crypt(u.PlainPassword)
	if err != nil{
		return err
	}
	u.Password = password

	return nil
}
