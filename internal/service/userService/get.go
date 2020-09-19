package userService

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
)

func Get(i *tracer.Infos, nickname string) (map[string]interface{}, error) {
	i.TraceIt("getting user")
	defer i.Span.Finish()

	uRepo := repository.NewUser()

	user, err := uRepo.GetUserByNick(i, nickname)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return user, nil
}
