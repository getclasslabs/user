package domains

import (
	"github.com/getclasslabs/user/internal/repository"
)

type User struct {
	Domain
	FirstName   string
	LastName    string
	BirthDate   string
	Gender      int
	Register    int
	Nickname    string
	Twitter     string
	Facebook    string
	Instagram   string
	Description string
	Telephone   string
	Address     string
}

func (u *User) Edit() error {
	uRepo := repository.NewUser()

	err := uRepo.Edit(u.Tracer, *u)

	if err != nil {
		u.Tracer.LogError(err)
		return err
	}
	return nil
}
