package userService

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/pkg"
	"github.com/getclasslabs/user/internal/repository"
	"github.com/getclasslabs/user/internal/service/login"
)

type ChangePassword struct{
	OldPassword string `json:"oldPass"`
	NewPassword string `json:"newPass"`
	NewPasswordConfirmation string `json:"newPassConfirmation"`
	Email string
}

func (c *ChangePassword) Validate(i *tracer.Infos) bool {
	i.TraceIt("validating")
	defer i.Span.Finish()

	l := login.Login{
		Email: c.Email,
		Password: c.OldPassword,
	}

	_, err := l.Validate(i)
	if err != nil{
		i.LogError(err)
		return false
	}
	return true
}

func (c *ChangePassword) Do(i *tracer.Infos) error{
	i.TraceIt("changing")
	defer i.Span.Finish()

	pass, err := pkg.Crypt(c.NewPassword)
	if err != nil {
		err = customerror.NewError(c, "crypting password", err)
		return err
	}

	uRepo := repository.NewUser()
	return uRepo.ChangePassword(i, c.Email, pass)

}