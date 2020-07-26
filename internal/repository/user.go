package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
)

const traceName = "user repository"

type User struct {
	db db.Database
}

func NewUser() *User {
	return &User{
		db: Db,
	}
}

func newMockedUser() *User {
	return &User{
		db: Mock{},
	}
}

func (u *User) SaveUser(i *tracer.Infos, email, password string) error {
	i.TraceIt(traceName)
	q := "INSERT INTO users(email,password) VALUES(?, ?) "
	_, err := u.db.Insert(i, q, email, password)

	if err != nil {
		err := customerror.NewDbError(u, q, err)
		i.LogError(err)
		return err
	}
	return nil
}