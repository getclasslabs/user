package login

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/pkg/kong"
	"github.com/getclasslabs/user/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Do(i *tracer.Infos) (map[string]interface{}, error) {
	i.TraceIt("login")

	result, err := l.Validate(i)
	if err != nil {
		return nil, err
	}

	k := kong.Service
	jwt, err := k.CreateCredentials(l.Email)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	result["jwt"] = jwt
	delete(result, "password")

	return result, nil
}

func (l *Login) Validate(i *tracer.Infos) (map[string]interface{}, error) {
	uRepo := repository.NewUser()
	result, err := uRepo.GetUserByEmail(i, l.Email)
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(l.Password))
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return result, nil
}
