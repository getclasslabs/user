package user

import (
	"github.com/getclasslabs/user/internal/pkg"
	"github.com/getclasslabs/user/internal/pkg/kong"
	"github.com/getclasslabs/user/internal/repository"
	"github.com/getclasslabs/user/tools"
)

type CreateUserService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUserService) Do(i *tools.Infos) (string, error) {

	i.Span = tools.TraceIt(i, "creating user")
	defer i.Span.Finish()

	pass := c.cryptPassword(c.Password)

	k := kong.Service

	err := k.CreateCustomer(c.Email)
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

func (c *CreateUserService) cryptPassword(plainPassword string) (string, error) {
	password, err := pkg.Crypt(plainPassword)
	if err != nil {
		return "", err
	}

	return password, nil
}
