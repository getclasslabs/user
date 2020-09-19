package domains

import (
	"github.com/getclasslabs/user/internal/repository"
)

const (
	StudentRegister = 0
	TeacherRegister = 1
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

	err := uRepo.Edit(u.Tracer,
		u.Email,
		u.Nickname,
		u.FirstName,
		u.LastName,
		u.BirthDate,
		u.Twitter,
		u.Facebook,
		u.Instagram,
		u.Description,
		u.Telephone,
		u.Address,
		u.Gender)

	if err != nil {
		u.Tracer.LogError(err)
		return err
	}
	return nil
}
