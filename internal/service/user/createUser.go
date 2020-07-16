package user

import (
	"github.com/getclasslabs/user/internal/domain"
	"github.com/getclasslabs/user/internal/pkg/kong"
	"github.com/getclasslabs/user/tools"
)

type Request struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}


func CreateUser(i *tools.Infos, userReq *Request) (string, error) {

	i.Span = tools.TraceIt(i, "creating user")
	defer i.Span.Finish()

	user, err := domain.NewUser(userReq.Email, userReq.Name, userReq.Password)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	k := kong.NewKong()

	err = k.CreateCustomer(user.Email)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	jwt, err := k.CreateCredentials(user.Email)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	return jwt, nil
}
