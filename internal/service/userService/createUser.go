package userService

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/pkg"
	"github.com/getclasslabs/user/internal/pkg/kong"
	"github.com/getclasslabs/user/internal/repository"
)

type CreateUserService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUserService) Do(i *tracer.Infos) (string, error) {

	i.TraceIt("creating user")
	defer i.Span.Finish()

	pass, err := pkg.Crypt(c.Password)
	if err != nil {
		err = customerror.NewError(c, "crypting password", err)
		return "", err
	}

	uRepo := repository.NewUser()

	err = uRepo.SaveUser(i, c.Email, pass)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	k := kong.Service

	err = k.CreateCustomer(c.Email)
	if err != nil {
		i.LogError(err)
		return "", err
	}

	jwt, err := k.CreateCredentials(c.Email)
	if err != nil {
		i.LogError(err)
		return "", err
	}

	return jwt, nil
}
