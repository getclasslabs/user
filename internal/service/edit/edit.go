package edit

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
	"github.com/getclasslabs/user/internal/service/user"
)

type Edit struct {
	user.Profile
	Twitter string `json:"twitter,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Description string `json:"description,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Address string `json:"address,omitempty"`

	//teacher
	Formation string `json:"formation,omitempty"`
	Specialization string `json:"specialization,omitempty"`
	WorkingTime int `json:"workingTime,omitempty"`
}

func (e *Edit) Do(i *tracer.Infos, email string) error {
	i.TraceIt("editing profile service")
	defer i.Span.Finish()

	urepo := repository.NewUser()

	result, err := urepo.GetUserByEmail(i, email)
	if err != nil {
		i.LogError(err)
		return err
	}

	if e.editCommonInfo(){

	}

	return nil
}


func (e *Edit) commonInfo() bool{

}


