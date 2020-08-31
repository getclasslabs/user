package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
)

const traceName = "user repository"

type User struct {
	db db.Database
	traceName string
}

func NewUser() *User {
	return &User{
		db: Db,
		traceName: "user repository",
	}
}

func newMockedUser() *User {
	return &User{
		db: Mock{},
	}
}

func (u *User) SaveUser(i *tracer.Infos, email, password string) error {
	i.TraceIt(u.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO users(email,password) VALUES(?, ?) "
	_, err := u.db.Insert(i, q, email, password)

	if err != nil {
		err := customerror.NewDbError(u, q, err)
		i.LogError(err)
		return err
	}
	return nil
}

func (u *User) SaveProfile(i *tracer.Infos, email string, register, gender int, firstName, lastName, birthDate, nickname string) error {
	i.TraceIt(traceName)
	defer i.Span.Finish()

	q := "UPDATE users SET " +
		"register=?, " +
		"gender=?, " +
		"first_name=?, " +
		"last_name=?, " +
		"birthDate=FROM_UNIXTIME(?), " +
		"nickname=? " +
		"WHERE " +
		"email = ?"
	_, err := u.db.Insert(i, q, register, gender, firstName, lastName, birthDate, nickname, email)

	if err != nil {
		err := customerror.NewDbError(u, q, err)
		i.LogError(err)
		return err
	}
	return nil
}

func (u *User) GetUserByEmail(i *tracer.Infos, email string) (map[string]interface{}, error) {
	i.TraceIt(traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"password, " +
		"nickname, " +
		"first_name, " +
		"last_name, " +
		"register, " +
		"gender " +
		"FROM users " +
		"WHERE " +
		"email = ?"

	result, err := u.db.Get(i, q, email)

	if err != nil {
		err := customerror.NewDbError(u, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil


}

